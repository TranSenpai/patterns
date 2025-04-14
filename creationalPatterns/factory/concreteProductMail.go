package factory

import "fmt"

// Concrete Products
type mail struct{}

func NewMail() *mail {
	return &mail{}
}

func (m mail) Send() {
	fmt.Println("Mail")
}
