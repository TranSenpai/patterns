package decorator

//  as known as wrapper

// - is a structural design pattern that lets you attach new behaviors to objects by
//   placing these objects inside special wrapper objects that contain the behaviors.

// When
// - scalability the behavior of an object

// How:
// - Create a common interface for all objects
// - Create concrete implementations of the interface
// - Create a decorator class that implements the same interface and has a reference to the object being decorated
// - Implement the same interface in the decorator class and delegate the calls to the wrapped object

import "fmt"

type NotificationServiceStrategy interface {
	Send(m string)
}

type sendMail struct{}

func (s sendMail) Send(m string) {
	fmt.Printf("Sending message: #{message} (Sender: Mail)\n")
}

type sendSMS struct{}

func (s sendSMS) Send(m string) {
	fmt.Printf("Sending message: #{message} (Sender: SMS)\n")
}

type NotificationService2 struct {
	notifier NotificationServiceStrategy
}

func (n NotificationService2) SendNotification(m string) {
	n.notifier.Send(m)
}

// Problem: Apply Strategy Pattern to send notification but how the application send both mail and SMS
type EmailSMSNotification struct {
	emailNotifier sendMail
	smsNotifier   sendSMS
}

// -> Nếu như thứ tự gọi quan trọng thì phải tạo ra 1 class mới để gọi cả 2 thứ
// -> Tăng thêm số lượng Notifiyer thì lại phải tăng thêm các struct cùng gọi:

func (n EmailSMSNotification) SendNotification(m string) {
	n.emailNotifier.Send(m)
	n.smsNotifier.Send(m)
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

// Decorator Pattern: is a structural design pattern that lets you attach new behaviors
// to objects by placing these objects inside special wrapper objects that contain the behaviors.

// Solution:
type NotifierDecorator struct {
	core     *NotifierDecorator
	notifier NotificationServiceStrategy
}

// Link to another decorator
// Its value is another notifier

// Wrapping the decorator with another decorator
// Use pointer and recursion to call the next decorator
// Work like stack
func (nd NotifierDecorator) Send(m string) {
	nd.notifier.Send(m)

	// Recursion: call the next decorator
	if nd.core != nil {
		nd.core.Send(m)
	}
}

func (nd NotifierDecorator) Decorate(notifier NotificationServiceStrategy) NotifierDecorator {
	return NotifierDecorator{
		// Địa chỉ decorator hiện tại
		core:     &nd,
		notifier: notifier,
	}
}

func NewNotifierDecorator(notifier NotificationServiceStrategy) NotifierDecorator {
	return NotifierDecorator{
		notifier: notifier,
	}
}

func Caller2() {
	notifier := NewNotifierDecorator(sendSMS{}).Decorate(sendMail{})
	s := NotificationService2{
		notifier: notifier,
	}
	s.SendNotification("Hello World")
}
