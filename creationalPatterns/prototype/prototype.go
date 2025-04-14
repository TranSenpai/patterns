package prototype

// as known as clone

// What
//	- is a creational design pattern
//  - lets you copy existing objects without making your code dependent on their classes

// Where
//	- delegates the cloning process to the actual objects that are being cloned
//	- declares a common interface for all objects that support cloning

// When
//  - have an object and want to create an exact copy of it

// Why
//  - to avoid have to go through all the fields of the original object and copy their
//    values over to the new object
//  - not all objects can be copied that way because some of the object’s fields
// 	  may be private and not visible from outside of the object itself
//  - the code depends on the class that have an object want to copy

// How
//	- create an interface that declares a clone method
// 	- implement the interface in the class that you want to clone
//  - create the prototype registry that will be used to store the prototypes

type Prototype interface {
	Clone() Prototype
}
