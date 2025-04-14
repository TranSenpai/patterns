package factory

import "fmt"

// Concrete Products
type app struct{}

func NewApp() *app {
	return &app{}
}

func (a app) Send() {
	fmt.Println("App")
}

type sms struct{}

func NewSMS() *sms {
	return &sms{}
}

func (s sms) Send() {
	fmt.Println("SMS")
}
