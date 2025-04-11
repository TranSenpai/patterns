package abstractfactory

import (
	"fmt"
)

// What
//	- is a creational design pattern
//	- produce families of related objects (các đối tượng liên quan)
// 	  without specifying their concrete classes (các lớp cụ thể)

// Where
//	- need a way to create individual objects so that they match other objects of the same family

// When
//	- don’t want to change existing code when adding new products or families of products to the program.

// Why
//	- avoid receive non-matching object
//	- easy scalability and maintaince

// How
// 	- explicitly declare interfaces for each distinct product of the product family
// 	- then make all variants (các loại) of products follow those interfaces.
// 	- declare the Abstract Factory—an interface with a list of creation methods
//    for all products that are part of the product family. These methods must return
//    abstract product types represented by the interfaces we extracted previously
//  - create a separate factory class based on the AbstractFactory interface.
// 	  A factory is a class that returns products of a particular kind.

type Drink interface {
	Drink()
}

type Food interface {
	Food()
}

type cake struct{}

func (c cake) Food() {
	fmt.Println("Cake")
}

type grilledOctopus struct{}

func (o grilledOctopus) Food() {
	fmt.Println("Grilled Octopus")
}

type beer struct{}

func (b beer) Drink() {
	fmt.Println("Beer")
}

type coffee struct{}

func (c coffee) Drink() {
	fmt.Println("Coffee")
}

// declare the Abstract Factory—an interface with a list of creation methods
// for all products that are part of the product family. These methods must return
// abstract product types represented by the interfaces we extracted previously
type ComboAbstractFactory interface {
	GetDrink() Drink
	GetFood() Food
}

type DrinkAndForgetTheWayHomeComboFactory struct{}

func (drunk DrinkAndForgetTheWayHomeComboFactory) GetDrink() Drink {
	return beer{}
}

func (drunk DrinkAndForgetTheWayHomeComboFactory) GetFood() Food {
	return grilledOctopus{}
}

type ComboName struct {
	lst map[string]ComboAbstractFactory
}

func (c ComboName) GetCombo(k string) ComboAbstractFactory {
	return c.lst[k]
}

func (c ComboName) SetCombo(k string, ca ComboAbstractFactory) {
	c.lst[k] = ca
}

func Caller() {
	// Create combo
	drunkCombo := DrinkAndForgetTheWayHomeComboFactory{}

	// Create combo list
	comboList := ComboName{
		lst: map[string]ComboAbstractFactory{},
	}

	// Add to combo list
	comboList.SetCombo("drunk", drunkCombo)

	// Get combo
	selectedCombo := comboList.GetCombo("drunk")

	// Get food and drink from combo
	selectedCombo.GetDrink().Drink() // Output: Beer
	selectedCombo.GetFood().Food()   // Output: Grilled Octopus
}
