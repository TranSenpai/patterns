package builder

// What
//	- is a creational design pattern
//	- allows creating complex objects step by step.
//	- separates the construction of a complex
//	  object from its representation
//	- allowing the same construction process
//	  to create different representations.

// Where
//	- extract the object con­struc­tion code out of its own
//	  class and move it to sep­a­rate objects called builders.

// When
//	- a com­plex object that requires labo­ri­ous (phức tạp),
//	  step-by-step ini­tial­iza­tion (quá trình khởi tạo)
// 	  of (bao gồm) many fields and nest­ed objects (các đối tượng lồng bên trong nhau).
//	- such ini­tial­iza­tion code is usu­al­ly buried inside
// 	  a mon­strous (khổng lồ) con­struc­tor with lots of para­me­ters.

// Why
//	- to avoid a long list of parameters in the constructor
//  - to avoid create a considerable number of subclasses
//  - easy to scalability representation of the object

// How
//	- declare a builder interface with product contruction steps are common to all types of builders
//    and methods getBuilder() to get type of builder
//  - declare  the
//	- declare a concrete builder that pro­vide dif­fer­ent imple­men­ta­tions of the con­struc­tion steps.
//	  Con­crete builders may pro­duce prod­ucts that don’t fol­low the com­mon interface.
//	- declare a director that is responsible for the construction process.
//    The director knows which builder to use and in which order to call the construction steps.

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloors()
	GetHouse() House
}
