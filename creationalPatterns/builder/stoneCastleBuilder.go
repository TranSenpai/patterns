package builder

type StoneCastleBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewStoneCastleBuilder() *StoneCastleBuilder {
	return &StoneCastleBuilder{}
}

func (b *StoneCastleBuilder) SetWindowType() {
	b.windowType = "Stone Window"
}

func (b *StoneCastleBuilder) SetDoorType() {
	b.doorType = "Iron and Stone Door"
}

func (b *StoneCastleBuilder) SetNumFloors() {
	b.floor = 10
}

func (b *StoneCastleBuilder) GetHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}
