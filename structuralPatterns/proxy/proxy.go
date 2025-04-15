package proxy

import (
	"fmt"
	"time"
)

// What
//	- is a structural design pattern
//	- lets you provide a substitute or placeholder for another object
//	- A proxy controls access to the original object,
//    allowing you to perform something either before or
//    after the request gets through to the original object

// When
//	- add logic before/after calling the original object without modifying it
//	- delaying the creation of resource-consuming objects until
//    they are actually needed (Lazy Initialization)

// Why
//	- save memory, avoid change original object

// How
//	- declare an interface for 2 origin class and proxy object
//	- concreate origin class and proxy class and implement the interface
//	- in the concreate proxy class, create the field interface type and use it

// Interface
type DataStorage interface {
	GetValue() int
}

// Origin class
type RealDataStorage struct{}

func (RealDataStorage) GetValue() int {
	// illustrate the time query db
	time.Sleep(time.Second * 5)
	return 100
}

// Proxy class
type ProxyDataStorage struct {
	cachedValue *int
	realStorage DataStorage
}

// cached value, if real data storgare get value too slow,
// return the cached value
func (s ProxyDataStorage) GetValue() int {
	if val := s.cachedValue; val != nil {
		return *val
	}
	val := s.realStorage.GetValue()
	s.cachedValue = &val
	return val
}

func NewProxyDataStorage(realStorage DataStorage) ProxyDataStorage {
	return ProxyDataStorage{realStorage: realStorage}
}

// Client use
type ValueService struct {
	storage DataStorage
}

func (s ValueService) FetchValue() int {
	return s.storage.GetValue()
}

func Caller() {
	value := ValueService{
		storage: NewProxyDataStorage(RealDataStorage{}),
	}.FetchValue()

	// First need to query
	fmt.Println(value)

	// Value is cached so it is faster
	fmt.Println(value)
}
