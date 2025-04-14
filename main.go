package main

import (
	"fmt"
	strategy "patterns/behaviorPatterns/strategy"
	abstractFactory "patterns/creationalPatterns/abstractFactory"
	builder "patterns/creationalPatterns/builder"
	factory "patterns/creationalPatterns/factory"
	prototype "patterns/creationalPatterns/prototype"
	singleton "patterns/creationalPatterns/singleton"
	adapter "patterns/structuralPatterns/adapter"
)

func main() {
	factory.Caller()
	strategy.Caller()
	singleton.Caller()
	abstractFactory.Caller()
	builder.Caller()
	prototype.Caller(20)
	adapter.Caller()
	fmt.Scanln()
}
