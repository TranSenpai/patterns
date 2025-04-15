package iterator

import (
	"fmt"
)

// What
// is a behavioral design pattern that lets you traverse elements of a collection
// without exposing its underlying representation (list, stack, tree, graph, etc.).

// When
//  Have many types of collections

// Problem
//  During the development of a software system, you may have many different collections of objects
//  that have many algorithms to traverse them, so when you add a new collection, add a new algorithm
//  -> Need an interface representing the tranverse algorithms

// How
// - create an interface that have getNext method representing the traverse algorithm and hasMore method to iterate
// - create a concrete class that implements the interface (implement the getNext method)
// - can create a collection of interators that have method createIterator new Interator interface

type Follower interface {
	Receiver(message string)
}

type Profile struct {
	name string
}

func (p Profile) Receiver(message string) {
	fmt.Printf("%s received message: %s\n", p.name, message)
}

// Interface for all interators
type FollowerIterator interface {
	next() Follower
	HasNext() bool
}

func sendMessage(iterator FollowerIterator, message string) {
	for iterator.HasNext() {
		iterator.next().Receiver(message)
	}
}

// Array: Interator
type FollowerArrayIterator struct {
	currentIdx int
	arr        []Follower
}

func NewFollowerArrayIterator(arr []Follower) *FollowerArrayIterator {
	return &FollowerArrayIterator{currentIdx: 0, arr: arr}
}

func (fi *FollowerArrayIterator) HasNext() bool {
	return len(fi.arr) > 0 && fi.currentIdx < len(fi.arr)
}

func (fi *FollowerArrayIterator) next() Follower {
	flw := fi.arr[fi.currentIdx]
	fi.currentIdx++
	return flw
}

// Linked List interator
type LinkedNode struct {
	val  Follower
	next *LinkedNode
}

type FollowerLinkedListIterator struct {
	node *LinkedNode
}

func NewFollowerLinkedListIterator(node *LinkedNode) *FollowerLinkedListIterator {
	return &FollowerLinkedListIterator{node: node}
}

func (fli *FollowerLinkedListIterator) HasNext() bool {
	return fli.node != nil
}

func (fli *FollowerLinkedListIterator) next() Follower {
	node := fli.node
	fli.node = node.next
	return node.val
}

var arrayOfFollower = []Follower{
	Profile{name: "John"},
	Profile{name: "Jane"},
	Profile{name: "Bob"},
	Profile{name: "Alice"},
}

var linkedListFollower = &LinkedNode{
	val: Profile{name: "John"},
	next: &LinkedNode{
		val: Profile{name: "Jane"},
		next: &LinkedNode{
			val: Profile{name: "Bob"},
			next: &LinkedNode{
				val:  Profile{name: "Alice"},
				next: nil,
			},
		},
	},
}

func Caller() {
	message := "hello"
	fmt.Println("a, b, c Array interator")
	var iterator FollowerIterator
	iterator = NewFollowerArrayIterator(arrayOfFollower)
	sendMessage(iterator, message)

	fmt.Println("a -> b -> c Linked List interator")
	iterator = NewFollowerLinkedListIterator(linkedListFollower)
	sendMessage(iterator, message)
}
