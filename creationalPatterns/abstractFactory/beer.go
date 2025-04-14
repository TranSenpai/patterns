package abstractfactory

import "fmt"

type beer struct{}

func (b beer) Drink() {
	fmt.Println("Beer")
}
