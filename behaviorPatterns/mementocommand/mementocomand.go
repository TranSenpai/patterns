package mementocommand

import (
	"errors"
	"fmt"
)

type Command interface {
	Execute()
	Undo()
}

type CommandAdd struct {
	v     *Value
	param int
}

type CommandSub struct {
	v     *Value
	param int
}

func (cmd CommandAdd) Execute() { cmd.v.Add(cmd.param) }
func (cmd CommandAdd) Undo()    { cmd.v.Sub(cmd.param) }

func (cmd CommandSub) Execute() { cmd.v.Sub(cmd.param) }
func (cmd CommandSub) Undo()    { cmd.v.Add(cmd.param) }

type Value struct {
	val int
}

func (v *Value) Add(num int) { v.val += num }
func (v *Value) Sub(num int) { v.val -= num }
func (v *Value) Val() int    { return v.val }

func NewValue(v int) Value {
	return Value{v}
}

type CommandNode struct {
	cmd  Command
	next *CommandNode
}

type CommandStack struct {
	current *CommandNode
}

func (cmdStack *CommandStack) Push(cmd Command) {
	cmdStack.current = &CommandNode{cmd: cmd, next: cmdStack.current}
}

func (cmdStack *CommandStack) Pop() (Command, error) {
	if cmdNode := cmdStack.current; cmdNode != nil {
		cmdStack.current = cmdNode.next
		return cmdNode.cmd, nil
	}
	return nil, errors.New("no command in stack")
}

type UndoableService struct {
	value    *Value
	addCmd   Command
	subCmd   Command
	cmdStack *CommandStack
}

func (s UndoableService) DoAdd() {
	s.addCmd.Execute()
	s.cmdStack.Push(s.addCmd)
}

func (s UndoableService) DoSub() {
	s.subCmd.Execute()
	s.cmdStack.Push(s.subCmd)
}

func (s UndoableService) GetValue() int { return s.value.val }

func (s UndoableService) Undo() {
	if cmd, err := s.cmdStack.Pop(); err == nil {
		cmd.Undo()
	}
}

func NewService(initValue int, incrStep int, decsStep int) UndoableService {
	value := NewValue(initValue)

	addCmd := CommandAdd{v: &value, param: incrStep}
	subCmd := CommandSub{v: &value, param: decsStep}

	return UndoableService{
		value:    &value,
		addCmd:   addCmd,
		subCmd:   subCmd,
		cmdStack: &CommandStack{},
	}
}

type ServiceMemento struct {
	vale     Value
	cmdStack *CommandStack
}

func (s *UndoableService) Save() ServiceMemento {
	return ServiceMemento{
		vale:     *s.value,
		cmdStack: &CommandStack{current: s.cmdStack.current},
	}
}

func (s *UndoableService) Restore(m ServiceMemento) {
	*s.value = m.vale
	s.cmdStack = m.cmdStack
}

type Caretaker struct {
	mementos []ServiceMemento
}

func (c *Caretaker) AddMemento(m ServiceMemento) {
	c.mementos = append(c.mementos, m)
}

func (c *Caretaker) Restore(s *UndoableService, idx int) {
	s.Restore(c.mementos[idx])
}

func (c *Caretaker) Size() int {
	return len(c.mementos)
}

func Caller() {
	careTaker := Caretaker{}

	service := NewService(10, 2, 1)
	fmt.Println(service.GetValue())

	service.DoAdd() // 12
	service.DoAdd() // 14
	fmt.Println(service.GetValue())

	// Save state idx 0
	careTaker.AddMemento(service.Save()) // 14

	service.DoSub() // 13
	fmt.Println(service.GetValue())

	// Save state idx 1
	careTaker.AddMemento(service.Save())

	service.Undo() // 14
	fmt.Println(service.GetValue())

	careTaker.AddMemento(service.Save())

	service.Undo() // 12
	service.Undo() // 10
	fmt.Println(service.GetValue())

	careTaker.AddMemento(service.Save())

	// Restore state
	for i := 0; i < careTaker.Size(); i++ {
		fmt.Println("Restored state at index: ", i)
		careTaker.Restore(&service, i)
		fmt.Println("Value: ", service.GetValue())

		service.Undo()
		fmt.Println("Value after undo: ", service.GetValue())
	}
}
