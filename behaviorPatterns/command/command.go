package command

import (
	"errors"
	"fmt"
)

// What
//	- is a behavioral design pattern that turns a request
//    into a stand-alone object that contains all information about the request.
//  - lets you pass requests as a method arguments, delay or queue a request’s execution
// 	  (có thể delay hoặc xếp nó vào hàng đợi để thực thi sau), and support undoable operations

// Why
//	- decouples sender and receiver, they depend on interface command
//	- queue them, delay execution, or batch them together
//	- since each action is an object, you can store them and call Undo() later

// How
//	- Create requester (invoker)
//	- Create an interface command that have the execute (and undo) method
//	- Create conrete command that specify what method of receicer use to solve problem
//	- Create receiver interface
//  - Create concrete receiver

// invoker
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

// interface command
type Command interface {
	Execute()
	Undo()
}

// conrete command
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

// Stack array to store for undo
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

// receiver
type Value struct {
	val int
}

func (v *Value) Add(num int) { v.val += num }
func (v *Value) Sub(num int) { v.val -= num }
func (v *Value) Val() int    { return v.val }

func NewValue(v int) Value {
	return Value{v}
}

func Caller() {
	service := NewService(10, 2, 1)
	fmt.Println(service.GetValue())

	service.DoAdd() // 12
	service.DoAdd() // 14
	fmt.Println(service.GetValue())

	service.DoSub() // 13
	fmt.Println(service.GetValue())

	service.Undo() // 14
	fmt.Println(service.GetValue())

	service.Undo() // 12
	service.Undo() // 10
	fmt.Println(service.GetValue())
}
