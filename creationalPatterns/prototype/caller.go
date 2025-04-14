package prototype

import (
	"fmt"
)

func Caller(numLoop int) {
	prototypeRegistry := &PrototypeRegistry{}
	p := Point{X: 1, Y: 2}
	for i := 0; i < numLoop; i++ {
		prototypeRegistry.AddItem(p.Clone())
	}

	for i := 0; i < len(prototypeRegistry.GetAllItems()); i++ {
		point := prototypeRegistry.GetItem(i)
		fmt.Printf("Cloned Point %d: %v\n", i+1, point)
	}

}
