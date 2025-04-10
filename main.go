package main

import (
	"fmt"
	f "patterns/factory"
	f2 "patterns/factory2"
)

func main() {

	factory := f.NewFactory()
	//
	sms := f.NewSMS()
	//
	factory.SetNotify("sms", sms)
	//
	factory.GetNotify("sms")
	factory.GetNotify("sms").Send()

	// fac2
	factory2 := f2.NewLogistic()
	truck := f2.NewTruck()
	factory2.SetLogistic("road", truck)
	factory2.GetLogistic("road")
	truck.Run()
	fmt.Scanln()
}

// ch := make(chan any)
// for i := 0; i < 30; i++ {
// 	go s.GetCarInstance(ch)
// 	a := <-ch
// 	fmt.Println(a)
// }

// for i := 0; i < 50; i++ {
// 	go s.GetEmployeeInstace(ch)
// 	a := <-ch
// 	fmt.Println(a)
// }
