package factory2

import "fmt"

type Logistic interface {
	Produce()
}

type factory2 struct {
	lst map[string]Transport
}

func NewLogistic() *factory2 {
	return &factory2{
		lst: map[string]Transport{},
	}
}

func (f2 *factory2) SetLogistic(k string, t Transport) {
	f2.lst[k] = t
}

func (f2 *factory2) GetLogistic(k string) Transport {
	return f2.lst[k]
}

type road struct{}

func (r *road) Produce() Transport {
	return NewTruck()
}

type sea struct{}

func (s *sea) Produce() Transport {
	return NewTruck()
}

type Transport interface {
	Run()
}

type truck struct{}

func (t truck) Run() {
	fmt.Println("Run")
}

func NewTruck() Transport {
	return &truck{}
}

type ship struct{}

func (s ship) Run() {
	fmt.Println("Ship Running")
}

func NewShip() Transport {
	return &ship{}
}
