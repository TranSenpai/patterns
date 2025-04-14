package builder

func Caller() {
	builderList := NewFactoryBuilder()

	builderList.RegisterBuilder("Wooden", NewWoodenHouseBuilder())
	builderList.RegisterBuilder("Stone", NewStoneCastleBuilder())

	director1 := NewDirector(builderList.GetBuilder("Wooden"))
	director1.ConstructHouse()

	director2 := NewDirector(builderList.GetBuilder("Stone"))
	director2.ConstructHouse()

	house := director1.builder.GetHouse()

	castle := director2.builder.GetHouse()

	println("House: ", house.windowType, house.doorType, house.floor)
	println("Castle: ", castle.windowType, castle.doorType, castle.floor)

}
