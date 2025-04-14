package abstractfactory

import "fmt"

type cake struct{}

func (c cake) Food() {
	fmt.Println("Cake")
}
