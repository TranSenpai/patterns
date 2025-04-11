package strategy

import "fmt"

// What?
// - Is a behavioral design pattern
// - Define a family of algorithms (các thuật toán tương đồng)
// - Put each of them (mỗi thuật toán) into a separate class
// - Make their objects interchangeable (hoán đổi)

// Where?
//	Dynamic switch between different behaviors or algorithms
//	Eliminate the long chains if-else or switch case statements

// When?
// 	Change behavior without change code
//	See a function

// Why?
// 	Easier to scalability and maintaint

// How?
// Take a class that does something specific (1 việc cụ thể) in a lot of different ways
// and extract all of these algorithms into separate classes called strategies.

// The original class, called context, must have a field for storing a
// reference (lưu trữ tham chiếu) to one of the strategies.

// The context delegates (ủy quyền) the work to a linked strategy
// object (đối tượng chiến lược) instead of executing it on its own.

// The context isn’t responsible for selecting an appropriate (phù hợp) algorithm for the job.

// False Example:
type NotificationService1 struct {
	notifierType string
}

func (s NotificationService1) Send(m string) {
	if s.notifierType == "mail" {
		fmt.Println("Send mail")
	} else if s.notifierType == "sms" {
		fmt.Println("Send SMS")
	} else {
		fmt.Println("False type")
	}
}

// True Example:

type NotificationServiceStrategy interface {
	Send(m string)
}

type sendMail struct{}

func (s sendMail) Send(m string) {
	fmt.Println("Send mail")
}

type sendSMS struct{}

func (s sendSMS) Send(m string) {
	fmt.Println("Send SMS")
}

type NotificationService2 struct {
	notifier NotificationServiceStrategy
}

func (n NotificationService2) SendNotification(m string) {
	n.notifier.Send(m)
}

func Caller() {
	sendMail := NotificationService2{
		notifier: sendMail{},
	}
	sendMail.SendNotification("Hello Mail")
	sendSMS := NotificationService2{
		notifier: sendSMS{},
	}
	sendSMS.SendNotification("Hello SMS")
}

// strategy adapt:
// Single Responsibility
// Open/Close
// Liskov subtituation
