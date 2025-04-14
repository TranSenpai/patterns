package adapter

func Print(hole RoundHole, adapter *SquarePegAdapter) {
	if hole.Fits(adapter) {
		println("Square peg fits in round hole")
	} else {
		println("Square peg does not fit in round hole")
	}
}

func Caller() {
	// Create a square peg
	sq := NewSquarePeg(10)

	// Use an adapter to take radius of the square peg
	SquarePegAdapter := SquarePegAdapter{peg: sq}

	// Create a round holes
	hole1 := NewRoundHole(5)
	hole2 := NewRoundHole(10)

	// Check if the square peg fits in the round hole
	Print(*hole1, &SquarePegAdapter)
	Print(*hole2, &SquarePegAdapter)

}
