package adapter

// Service
type SquarePeg struct {
	width int
}

func NewSquarePeg(width int) *SquarePeg {
	return &SquarePeg{width: width}
}

func (s *SquarePeg) GetWidth() int {
	return s.width
}
