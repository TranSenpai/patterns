package adapter

// as known as wrapper pattern

import (
	"math"
)

// What
//  - is a struc­tur­al design pat­tern
//  - allows objects with incom­pat­i­ble (không phù hợp) inter­faces to collaborate (làm việc với nhau).

// Where
//  - is used when you want to use an ex­ist­ing class, but its in­ter­face does not match the one you need.

// When
//  - when want two in­com­pat­i­ble ob­jects to work to­gether

// How
//  - create an adapter class that wraps the in­com­pat­i­ble ob­ject and
//    provides a new all methods that is com­pat­i­ble with the one you need.

// Adapter
type SquarePegAdapter struct {
	peg *SquarePeg
}

func (a *SquarePegAdapter) GetRadius() int {
	return int(float64(a.peg.GetWidth()) * math.Sqrt(2) / 2) // Example conversion logic
}
