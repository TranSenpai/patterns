package singleton

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
