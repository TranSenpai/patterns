package main

import (
	"fmt"
	chain "patterns/behaviorPatterns/chainOfResponsibility"
	command "patterns/behaviorPatterns/command"
	iterator "patterns/behaviorPatterns/iterator"
	mediator "patterns/behaviorPatterns/mediator"
	memento "patterns/behaviorPatterns/memento"
	mementocommand "patterns/behaviorPatterns/mementocommand"
	observer "patterns/behaviorPatterns/observer"
	state "patterns/behaviorPatterns/state"
	strategy "patterns/behaviorPatterns/strategy"
	templatemethod "patterns/behaviorPatterns/templatemethod"
	abstractFactory "patterns/creationalPatterns/abstractFactory"
	builder "patterns/creationalPatterns/builder"
	factory "patterns/creationalPatterns/factory"
	prototype "patterns/creationalPatterns/prototype"
	singleton "patterns/creationalPatterns/singleton"
	adapter "patterns/structuralPatterns/adapter"
	bridge "patterns/structuralPatterns/bridge"
	composite "patterns/structuralPatterns/composite"
	decorator "patterns/structuralPatterns/decorator"
	facade "patterns/structuralPatterns/facade"
	flyweight "patterns/structuralPatterns/flyweight"
	proxy "patterns/structuralPatterns/proxy"
)

func main() {
	factory.Caller()
	strategy.Caller()
	singleton.Caller()
	abstractFactory.Caller()
	builder.Caller()
	prototype.Caller(20)
	adapter.Caller()
	bridge.Caller()
	flyweight.Caller()
	composite.Caller()
	decorator.Caller2()
	chain.Caller()
	facade.Caller()
	iterator.Caller()
	templatemethod.Caller()
	proxy.Caller()
	observer.Caller()
	state.Caller()
	mediator.Caller()
	command.Caller()
	memento.Caller()
	mementocommand.Caller()
	fmt.Scanln()
}
