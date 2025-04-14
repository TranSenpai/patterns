package abstractfactory

import "fmt"

type coffee struct{}

func (c coffee) Drink() {
	fmt.Println("Coffee")
}
