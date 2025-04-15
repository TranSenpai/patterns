package mediator

import (
	"fmt"
	"sync"
)

// as known as intermediary, controller

// What
//	- is a behavioral design pattern that lets you reduce
//	  chaotic dependencies between objects. The pattern restricts
// 	  direct communications between the objects and forces them to
// 	  collaborate (làm việc) only via a mediator object (thông qua controller).

// Why
//	- control the work flow between many class
//	- giảm phụ thuộc giữa các class giao tiếp với nhau mà hướng nó tới 1 interface giao tiếp

// How
//	- define interface that declares methods for communication between components
//	- create concrete mediator
//	- create component interface
//	- create concrete component

// Shared resource
type Tool struct {
	isUsing bool
}

func (t *Tool) CanUse() bool { return !t.isUsing }
func (t *Tool) Using()       { t.isUsing = true }
func (t *Tool) Done()        { t.isUsing = false }

// Interface mediator
type Mediator interface {
	RegisterUsingTool(c Component)
	NotifyDone(fromComp Component)
	Wait()
}

// Implement mediator
type SimpleMediator struct {
	tool      *Tool
	queue     []Component
	lock      sync.Mutex
	waitgroup sync.WaitGroup
}

func (m *SimpleMediator) allocToolForComponent(tool *Tool, c Component) {
	tool.Using()
	c.AllowUsingTool(m.tool)
}

func (m *SimpleMediator) RegisterUsingTool(c Component) {
	m.waitgroup.Add(1)
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.tool.CanUse() {
		m.allocToolForComponent(m.tool, c)
		return
	}
	m.queue = append(m.queue, c)
}

func (m *SimpleMediator) NotifyDone(Component) {
	m.waitgroup.Done()
	m.lock.Lock()
	defer m.lock.Unlock()
	fmt.Println("component done its job")
	m.tool.Done()

	if len(m.queue) == 0 {
		return
	}

	nextCom := m.queue[0]
	m.queue = append(make([]Component, 0), m.queue[1:]...)

	m.allocToolForComponent(m.tool, nextCom)
}

func (m *SimpleMediator) Wait() {
	m.waitgroup.Wait()
}

func NewMediator(tool *Tool) Mediator {
	return &SimpleMediator{
		tool:      tool,
		lock:      sync.Mutex{},
		waitgroup: sync.WaitGroup{},
	}
}

// interface object
type Component interface {
	AllowUsingTool(tool *Tool)
	Solve()
}

// Object that will communicate via mediator
type Worker struct {
	// the channel they communicate
	mediator Mediator
}

func (w Worker) doJob(tool *Tool) {
	fmt.Println("Worker is using tool")
	w.mediator.NotifyDone(w)
}

func (w Worker) Solve() {
	fmt.Println("Worker needs to use tool to solve problem. He is asking mediator...")
	w.mediator.RegisterUsingTool(w)
}

func (w Worker) AllowUsingTool(tool *Tool) {
	go w.doJob(tool)
}

type Engineer struct {
	mediator Mediator
}

func (e Engineer) design(tool *Tool) {
	fmt.Println("Engineer is using tool")
	e.mediator.NotifyDone(e)
}

func (e Engineer) Solve() {
	fmt.Println("Engineer needs to use tool to solve problem. He is asking mediator...")
	e.mediator.RegisterUsingTool(e)
}

func (e Engineer) AllowUsingTool(tool *Tool) {
	go e.design(tool)
}

func Caller() {
	sharingTool := &Tool{isUsing: false}
	mediator := &SimpleMediator{tool: sharingTool}

	components := []Component{
		Worker{mediator: mediator},
		Worker{mediator: mediator},
		Worker{mediator: mediator},
		Engineer{mediator: mediator},
		Engineer{mediator: mediator},
	}

	for _, c := range components {
		c.Solve()
	}

	mediator.Wait()
	fmt.Println("All components have done their jobs")
}
