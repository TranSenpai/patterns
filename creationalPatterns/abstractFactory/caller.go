package abstractfactory

func Caller() {
	// Create combo
	drunkCombo := DrinkAndForgetTheWayHomeComboFactory{}

	// Create combo list
	comboList1 := ComboName{
		lst: map[string]ComboAbstractFactory{},
	}

	// Add to combo list
	comboList1.SetCombo("drunk", drunkCombo)

	// Get combo
	selectedCombo1 := comboList1.GetCombo("drunk")

	// Create combo
	healthyCombo := MorningHealthy{}

	// Create combo list
	comboList2 := ComboName{
		lst: map[string]ComboAbstractFactory{},
	}

	// Add to combo list
	comboList2.SetCombo("healthy", healthyCombo)

	// Get combo
	selectedCombo2 := comboList2.GetCombo("healthy")

	// Get food and drink from combo
	selectedCombo1.GetDrink().Drink() // Output: Beer
	selectedCombo1.GetFood().Food()   // Output: Grilled Octopus

	selectedCombo2.GetDrink().Drink() // Output: Coffee
	selectedCombo2.GetFood().Food()   // Output: Cake
}
