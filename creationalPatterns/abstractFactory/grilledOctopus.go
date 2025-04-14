package abstractfactory

import "fmt"

type grilledOctopus struct{}

func (o grilledOctopus) Food() {
	fmt.Println("Grilled Octopus")
}
