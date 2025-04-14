package factory

// as known as Virtual Constructor
// What:
//	- is a creational design pattern
//	- provides an interface for creating objects in a superclass
//	- but allows subclasses to alter (thay đổi) the type of objects that will be created

// Where:
//	- need to create objects but don’t want to specify their concrete (cụ thể) classes
//	- want the client code to work with interfaces or abstract classes rather than specific implementations
//	- foresee adding more types of objects in the future

// When:
//	- object types may vary at runtime or in future
//	- want to decouple the instantiation logic from usage

// Why:
//	- easy to maintance code and scalability
//	- client just interacts with abstract class or interface

// How:
//	- declares the Product interface, which is common (phổ biến) to all
// 	  objects that can be produced by the creator and its subclasses

//	- concrete Products are different implementations of the product interface.

//	- create a struct (class) has a factory method or just factory method
//	  that return's type of this method matches the product interface

// Factory struct (class) have factory method
// Or just a factory method

type notifyFactory struct {
	lst map[string]Notify
}

func NewFactory() *notifyFactory {
	return &notifyFactory{
		lst: map[string]Notify{},
	}
}

func (n notifyFactory) GetNotify(kind string) Notify {
	return n.lst[kind]
}

func (n notifyFactory) SetNotify(kind string, not Notify) {
	n.lst[kind] = not
}

// factory adapt
