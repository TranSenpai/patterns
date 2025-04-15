package visitor

import "fmt"

// What
//   - is a behavioral design pattern that lets you separate algorithms
//     from the objects on which (mà chúng) they operate.

// Why
//   - easy to add new algorithm without rewrite old source

// How
//   - Create an interface that have all common method of subclass and the accept method
//   - In subclasses, implement the interface
//   - Create a visitor interface that have all method visit all types object but return the general type
//   - New logic with implement visitor interface to interact with subclasses
type Visitor interface {
	VisitProduct(product Product)
	VisitCategory(category Catelory)
	VisitUser(user User)
}

type PrinterVisitor struct{}

func (PrinterVisitor) VisitProduct(product Product) {
	fmt.Printf("Visiting product: %v\n", product)
}

func (PrinterVisitor) VisitCategory(category Catelory) {
	fmt.Printf("Visiting category: %v\n", category)
}

func (PrinterVisitor) VisitUser(user User) {
	fmt.Printf("Visiting user: %v\n", user)
}

type Visitable interface {
	Accept(visitor Visitor)
}

type Catelory struct {
	Title string
}
type Product struct {
	Name string
}
type User struct {
	FirstName string
	LastName  string
}

func (c Catelory) Accept(visitor Visitor) {
	visitor.VisitCategory(c)
}

func (p Product) Accept(visitor Visitor) {
	visitor.VisitProduct(p)
}

func (u User) Accept(visitor Visitor) {
	visitor.VisitUser(u)
}

// another algorithm
type JSONEncoderVisitor struct{}

func (JSONEncoderVisitor) VisitProduct(product Product) {
	fmt.Printf("Visiting product: %v\n", product)
}

func (JSONEncoderVisitor) VisitCategory(category Catelory) {
	fmt.Printf("Visiting category: %v\n", category)
}

func (JSONEncoderVisitor) VisitUser(user User) {
	fmt.Printf("Visiting user: %v\n", user)
}

func Caller() {
	items := []Visitable{
		Catelory{Title: "Electronics"},
		Product{Name: "iPhone"},
		User{FirstName: "John", LastName: "Doe"},
	}

	printerVisitor := PrinterVisitor{}
	jsonEncoderVisitor := JSONEncoderVisitor{}

	for i := range items {
		items[i].Accept(printerVisitor)
		items[i].Accept(jsonEncoderVisitor)
	}
}
