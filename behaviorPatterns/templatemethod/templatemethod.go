package templatemethod

import "fmt"

// What
//	- is a behavioral design pattern that defines the skeleton (bộ khung) of an
// 	  algorithm in the superclass but lets subclasses override specific
//    steps of the algorithm without changing its structure.

// Where
//	- in multi-layered systems there is a common process but the specific
//    implementation of each step may differ.
//	- have a process (algorithm) divided into multiple steps,
//    where some steps are fixed, some can be changed by subclasses.

// When
//	- want to reuse the structural algorithm between multiple layers, changing the details

// Why
//	- Easey to understand the logic, avoid duplicating process code

type WebPage interface {
	RenderHeader()
	RenderBody()
	RenderFooter()
}

type WebTemplate struct{}

func (WebTemplate) RenderHeader() { fmt.Println("This is header") }

func (WebTemplate) RenderFooter() { fmt.Println("This is footer") }

type HomePage struct {
	WebTemplate
}

func (HomePage) RenderBody() { fmt.Println("This is body") }

type ProductPage struct {
	WebTemplate
}

func (ProductPage) RenderBody() { fmt.Println("This is body") }

func RenderWebsite(website WebPage) {
	website.RenderHeader()
	website.RenderBody()
	website.RenderFooter()
}

func Caller() {
	homepage := HomePage{}
	productpage := ProductPage{}

	RenderWebsite(homepage)
	RenderWebsite(productpage)
}
