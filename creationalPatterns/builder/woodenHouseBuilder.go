package builder

type WoodenHouseBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewWoodenHouseBuilder() *WoodenHouseBuilder {
	return &WoodenHouseBuilder{}
}

func (b *WoodenHouseBuilder) SetWindowType() {
	b.windowType = "Wooden Window"
}

func (b *WoodenHouseBuilder) SetDoorType() {
	b.doorType = "Wooden Door"
}

func (b *WoodenHouseBuilder) SetNumFloors() {
	b.floor = 2
}

func (b *WoodenHouseBuilder) GetHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}
