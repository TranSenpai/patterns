package factory

import "fmt"

type Notify interface {
	Send()
}

type mail struct {
}

func NewMail() *mail {
	return &mail{}
}

func (m mail) Send() {
	fmt.Println("Mail")
}

type app struct {
}

func NewApp() *app {
	return &app{}
}

func (a app) Send() {
	fmt.Println("App")
}

type sms struct {
}

func NewSMS() *sms {
	return &sms{}
}

func (s sms) Send() {
	fmt.Println("SMS")
}
