package memento

import "fmt"

// What
//	- is a behavioral design pattern that lets you save and restore the previous
// 	- state of an object without revealing (tiết lộ) the details of its implementation.

// When
//	- want to restore a previous state of an object
//	- encapsulate internal state without violating encapsulation
//	- don’t want external objects to be able to directly modify the internals (các thành phần)
//    of another object but still want to preserve/restore states.

// How
//	- create an origin class
//	- declare a snapShot and restore function in origin class that return Memento type
//	- create a memento class that have all fields of origin class
//	- create a Caretakers to store many Memento (add, restore mementory[index] function)

// Originator and Mementor have to be in the same package

// Origin
type Editor struct {
	str string
}

func (e *Editor) TypeMore(s string) { e.str += s }
func (e *Editor) Content() string   { return e.str }

// declare a snapShot and restore function
func (e *Editor) Save() Memento     { return Memento{content: e.str} }
func (e *Editor) Restore(m Memento) { e.str = m.content }

type Memento struct {
	// Because memento copy origin
	// Memento must declare all origin's field
	content string
}

type Caretaker struct {
	mementos []Memento
}

func (c *Caretaker) Add(m Memento) {
	c.mementos = append(c.mementos, m)
}

func (c *Caretaker) Restore(e *Editor, idx int) {
	e.Restore(c.mementos[idx])
}

func (c *Caretaker) Size() int {
	return len(c.mementos)
}

func NewEditor(s string) Editor {
	return Editor{s}
}

func Caller() {
	editor := NewEditor("")
	caretaker := &Caretaker{}

	editor.TypeMore("hellow")
	editor.TypeMore("world")

	caretaker.Add(editor.Save())

	editor.TypeMore("I am Memento")
	caretaker.Add(editor.Save())

	editor.TypeMore(", a design pattern lets you save and")
	editor.TypeMore("restore the previous state")

	for i := 0; i < caretaker.Size(); i++ {
		caretaker.Restore(&editor, i)
		fmt.Println("Restored content: ", editor.Content())
	}
}
