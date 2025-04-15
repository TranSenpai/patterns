package state

import (
	"errors"
	"fmt"
	"log"
)

// What
//	- is a behavioral design pattern that lets an object
//    alter its behavior when its internal state changes.
//    It appears as if the object changed its class.

// When
//	- changes depending on its internal state
//	- has multiple states, and each state triggers different behaviors for the same method call
//	- avoid large if-else / switch-case statements based on state

// Why
//  - avoids complex if-else
//  - easier to add new states
//	- each state is its own class (struct), following Single Responsibility Principle

// How
//  - define a State interface that each state will implement
//	- create concrete state struct
//	- create context class

var ErrInvalidAction = errors.New("invalid action")

type OrderState interface {
	Cancel() error
	Pay() error
	Deliver() error
	Finish() error
	String() string
}

type Order struct {
	state OrderState
}

type OrderStateCancelled struct {
	order *Order
}

func (OrderStateCancelled) Cancel() error  { return ErrInvalidAction }
func (OrderStateCancelled) Pay() error     { return ErrInvalidAction }
func (OrderStateCancelled) Deliver() error { return ErrInvalidAction }
func (OrderStateCancelled) Finish() error  { return ErrInvalidAction }
func (OrderStateCancelled) String() string { return "cancel" }

type OrderStateFinish struct {
	order *Order
}

func (OrderStateFinish) Cancel() error  { return ErrInvalidAction }
func (OrderStateFinish) Pay() error     { return ErrInvalidAction }
func (OrderStateFinish) Deliver() error { return ErrInvalidAction }
func (OrderStateFinish) Finish() error  { return ErrInvalidAction }
func (OrderStateFinish) String() string { return "finished" }

type OrderStateCreated struct {
	order *Order
}

func (state OrderStateCreated) Cancel() error {
	state.order.updateSate(OrderStateCancelled{order: state.order})
	return nil
}

func (state OrderStateCreated) Pay() error {
	state.order.updateSate(OrderStatePaid{order: state.order})
	return nil
}

func (OrderStateCreated) Deliver() error { return ErrInvalidAction }
func (OrderStateCreated) Finish() error  { return ErrInvalidAction }
func (OrderStateCreated) String() string { return "created" }

type OrderStatePaid struct {
	order *Order
}

func (state OrderStatePaid) Deliver() error {
	state.order.updateSate(OrderStateDelivered{order: state.order})
	return nil
}

func (OrderStatePaid) Cancel() error  { return ErrInvalidAction }
func (OrderStatePaid) Pay() error     { return ErrInvalidAction }
func (OrderStatePaid) Finish() error  { return ErrInvalidAction }
func (OrderStatePaid) String() string { return "paid" }

type OrderStateDelivered struct {
	order *Order
}

func (state OrderStateDelivered) Finish() error {
	state.order.updateSate(OrderStateFinish{order: state.order})
	return nil
}

func (OrderStateDelivered) Deliver() error { return ErrInvalidAction }
func (OrderStateDelivered) Cancel() error  { return ErrInvalidAction }
func (OrderStateDelivered) Pay() error     { return ErrInvalidAction }
func (OrderStateDelivered) String() string { return "delivered" }

func (o *Order) updateSate(state OrderState) {
	log.Printf("order has changed state to: %v\n", o.state)
}

func NewOrder() *Order {
	order := &Order{}
	initState := OrderStateCreated{order: order}
	order.state = initState

	return order
}

func (o *Order) CurrentState() OrderState { return o.state }

func (o *Order) Cancel() error { return o.state.Cancel() }

func (o *Order) Pay() error { return o.state.Pay() }

func (o *Order) Deliver() error { return o.state.Deliver() }

func (o *Order) Finish() error { return o.state.Finish() }

func Caller() {
	order := NewOrder()

	fmt.Println("Order state: ", order.CurrentState())

	if err := order.Finish(); err != nil {
		log.Println(err)
	}

	if err := order.Pay(); err != nil {
		log.Println(err)
	}

	if err := order.Deliver(); err != nil {
		log.Println(err)
	}

	if err := order.Finish(); err != nil {
		log.Println(err)
	}
}
