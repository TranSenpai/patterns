package singleton

// What
//	- is a creational design pattern
//	- ensure a class has only one instance
//	- provide a global point of access to it

// Where
//	- when you want to control the number of instances
//	  and provide a single point of access to it

// When
//	- want to make sure that a class has only one instance

// Why
// 	- to control access to shared resources, such as a file or a database connection
//	- to provide a global point of access to the instance
//	- to avoid the overhead of creating multiple instances of a class
//	- to provide a single point of access to the instance

// How
//	- use a private constructor to prevent instantiation from outside the class
//	- use a method to provide access to the instance
//	- use a global variable to hold the instance
//	- use a mutex to ensure thread safety

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type car struct {
	name  string
	brand string
}

var singleCarInstance *car

func GetCarInstance(ch chan any) {
	// check instance created or not
	if singleCarInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		// because many goroutines can pass the first check
		// the second check make sure if the first goroutine
		// created the instance, orther can not create again
		if singleCarInstance == nil {
			fmt.Println("Create car instance now.")
			singleCarInstance = &car{}
		} else {
			fmt.Println("Car instance has created.")
		}
	} else {
		fmt.Println("Car instance has created.")
	}
	ch <- &singleCarInstance
}

var once sync.Once

type employee struct {
	name   string
	career string
	id     int
}

var singleEmployeeInstace *employee

func GetEmployeeInstace(ch chan any) {
	if singleEmployeeInstace == nil {
		// once.Do in sync package take func parameter
		// and make sure it just run once time
		once.Do(
			func() {
				fmt.Println("Create Employee instance now.")
				singleEmployeeInstace = &employee{}
			})
	} else {
		fmt.Println("Employee instace has created")
	}
	ch <- &singleEmployeeInstace
}

func Caller() {
	ch := make(chan any)
	for i := 0; i < 3; i++ {
		go GetCarInstance(ch)
		a := <-ch
		fmt.Println(a)
	}

	for i := 0; i < 3; i++ {
		go GetEmployeeInstace(ch)
		a := <-ch
		fmt.Println(a)
	}
}

// Singleton adapt:
