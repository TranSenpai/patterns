package prototype

type Point struct {
	X int
	Y int
}

func (p *Point) Clone() Prototype {
	return &Point{X: p.X, Y: p.Y}
}
