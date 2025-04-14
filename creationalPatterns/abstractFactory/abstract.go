package abstractfactory

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

// declare the Abstract Factory—an interface with a list of creation methods
// for all products that are part of the product family. These methods must return
// abstract product types represented by the interfaces we extracted previously
type ComboAbstractFactory interface {
	GetDrink() Drink
	GetFood() Food
}

// Family 1:
type DrinkAndForgetTheWayHomeComboFactory struct{}

func (drunk DrinkAndForgetTheWayHomeComboFactory) GetDrink() Drink {
	return beer{}
}

func (drunk DrinkAndForgetTheWayHomeComboFactory) GetFood() Food {
	return grilledOctopus{}
}

// Family 2:

type MorningHealthy struct{}

func (m MorningHealthy) GetDrink() Drink {
	return coffee{}
}

func (m MorningHealthy) GetFood() Food {
	return cake{}
}
