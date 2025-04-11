package main

import (
	"fmt"
	strategy "patterns/behaviorPatterns/strategy"
	abstractFactory "patterns/creationalPatterns/abstractFactory"
	factory "patterns/creationalPatterns/factory"
	singleton "patterns/creationalPatterns/singleton"
)

func main() {
	factory.Caller()
	strategy.Caller()
	singleton.Caller()
	abstractFactory.Caller()
	fmt.Scanln()
}
