package observer

import (
	"fmt"
)

// as known as event-subscriber, listener

// What
//	- is a behavioral design pattern that lets you define
//    a subscription mechanism (cơ chế đăng ký) to notify multiple objects
//    about any events that happen to the object they’re observing.

// When
//	- want observer object update or get notification when a subject changes

// Why
//	- make sure that observer and subject relationship is loose relationship

type Job struct {
	Title string
}

type Observer interface {
	ReceiveNotify(j Job)
}

// Subcriber
type Developer struct{}

func (Developer) ReceiveNotify(j Job) { fmt.Println("Many thank, I 've received job:", j.Title) }

// Publisher
type ITJobsCompany struct {
	jobs      []Job
	observers []Observer
}

func (comp *ITJobsCompany) AddObserver(o Observer) {
	comp.observers = append(comp.observers, o)
}

func (comp *ITJobsCompany) RemoveObserver(o Observer) {
	for i := range comp.observers {
		if comp.observers[i] == o {
			comp.observers = append(comp.observers[:i], comp.observers[i+1:]...)
			return
		}
	}
}

func (comp *ITJobsCompany) notifyToObservers(j Job) {
	for i := range comp.observers {
		comp.observers[i].ReceiveNotify(j)
	}
}

func (comp *ITJobsCompany) AddNewJob(j Job) {
	comp.jobs = append(comp.jobs, j)
	comp.notifyToObservers(j)
}

func Caller() {
	itComp := ITJobsCompany{}
	dev := Developer{}

	itComp.AddObserver(dev)

	job1 := Job{Title: "Fresher developer"}
	job2 := Job{Title: "Junior developer"}
	itComp.AddNewJob(job1)
	itComp.AddNewJob(job2)

	itComp.RemoveObserver(dev)

	job3 := Job{Title: "Senior developer"}
	itComp.AddNewJob(job3)
}
