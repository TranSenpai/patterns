package adapter

// Client
type RoundHole struct {
	radius int
}

func NewRoundHole(radius int) *RoundHole {
	return &RoundHole{radius: radius}
}

func (h *RoundHole) GetRadius() int {
	return h.radius
}

func (h *RoundHole) Fits(peg IRoundPeg) bool {
	return peg.GetRadius() <= h.radius
}
