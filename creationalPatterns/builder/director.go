package builder

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{builder: b}
}

func (d *Director) ConstructHouse() {
	d.builder.SetWindowType()
	d.builder.SetDoorType()
	d.builder.SetNumFloors()
}
