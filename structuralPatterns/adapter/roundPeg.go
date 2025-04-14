package adapter

type IRoundPeg interface {
	GetRadius() int
}

type RoundPeg struct {
	radius int
}

func NewRoundPeg(radius int) *RoundPeg {
	return &RoundPeg{}
}

func (r *RoundPeg) GetRadius() int {
	return r.radius
}
