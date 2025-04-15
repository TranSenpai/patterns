package composite

// as known as object tree

// Composite is a structural design pattern that lets you compose (kết hợp)
// objects into (thành) tree structures and then work with these structures
// as if they were individual objects (các đối tượng riêng lẻ).

// -> Answer the question: "how to handle (xử lý) individual (riêng lẻ) objects
//    and compositions of objects uniformly (cùng 1 cách)?"
// -> use when: you want to represent a part-whole hierarchy (cấu trúc phân cấp phần và toàn bộ)

// How:
// Define a common interface for all objects in the tree structure (cấu trúc cây)
// Define leaf objects (các đối tượng lá) that implement the common interface
// Define composite objects (các đối tượng tổng hợp) that also implement the common interface

// Problem:

type Item1 struct {
	Name  string
	Price float32
	// Ban đầu không có children nhưng bởi vì nghiệp vụ về sau nên phải bổ sung
	children []Item1
}

func (item Item1) cost1() float32 {
	cost := item.Price
	for _, child := range item.children {
		cost += child.cost1()
	}
	return cost + item.Price
}

func CreatePackage1() Item1 {
	return Item1{
		// A box is not a item
		// So it dont have price
		Name:  "root box",
		Price: 0,
		children: []Item1{
			{
				Name:     "Mouse",
				Price:    10,
				children: nil,
			},
			{
				Name:  "Sub box",
				Price: 20,
				children: []Item1{
					{
						Name:  "Sub box 2",
						Price: 0,
						children: []Item1{
							{
								Name:  "Keyboard",
								Price: 30,
							},
							{
								Name:  "Monitor",
								Price: 0,
							},
						},
					},
				},
			},
		},
	}
}

// Solution:

// Generalize the interface for all objects in the tree structure
type Item interface {
	Cost() float32
}

// Define leaf objects that implement the common interface

type RealItem struct {
	Name  string
	Price float32
}

func (item RealItem) Cost() float32 {
	return item.Price
}

// Define composite objects that also implement the common interface
type Box struct {
	children []Item
}

func (box Box) Cost() float32 {
	var cost float32 = 0.0

	for _, item := range box.children {
		cost += item.Cost()
	}
	return cost
}

// Use
func CreatePackage() Item {
	return Box{
		children: []Item{
			RealItem{Name: "Mouse", Price: 10},
			Box{
				children: []Item{
					Box{
						children: []Item{
							RealItem{Name: "Keyboard", Price: 30},
							RealItem{Name: "Monitor", Price: 0},
						},
					},
				},
			},
		},
	}
}

func Caller() {
	// Create a package with items
	package1 := CreatePackage()
	package2 := CreatePackage1()

	// Calculate the total cost of the package using the composite pattern
	totalCost := package1.Cost()
	totalCost1 := package2.cost1()

	println(totalCost)  // Output: 40.0
	println(totalCost1) // Output: 40.0
}

// The composite pattern allows you to treat individual objects and compositions of objects uniformly.
// This is useful when you want to represent a part-whole hierarchy, such as a file system or a graphical scene graph.
