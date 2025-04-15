package flyweight

// as known as cache

import "fmt"

// What
//  - is a struc­tur­al design pat­tern
//  - lets you fit more objects into the available amount of RAM
//    by sharing common parts of state between multiple objects
//    instead of keeping all of the data in each object

// Not optizime memory when not using flyweight pattern

type NotOptimizeChatMessage struct {
	Content      string
	Sender       string
	SenderAvatar []byte
}

func NotOptimizeChat() {
	fmt.Println([]NotOptimizeChatMessage{
		{
			Content:      "Hello",
			Sender:       "Alice",
			SenderAvatar: make([]byte, 1024*300), // 300kb
		},
		{
			Content:      "Hi",
			Sender:       "Bob",
			SenderAvatar: make([]byte, 1024*400), // 400kb
		},
		{
			Content:      "How are you doing?",
			Sender:       "Alice",
			SenderAvatar: make([]byte, 1024*300), // 300kb
		},
		{
			Content:      "I am fine",
			Sender:       "Bob",
			SenderAvatar: make([]byte, 1024*400), // 400kb
		},
	})
	// Everytime Alice and Bob send a message,
	// a lot of memory being wasted to represent their avatars.

	// Total memory usage: 300 + 400 + 300 + 400 = 1400kb = 1.4mb
	// -> Not optimize memory usage
	// -> Need to reuse memory for the same object (in this situation, the same avatar)
}

type ChatMessage struct {
	Content string
	Sender  *Sender
}

type Sender struct {
	Name   string
	Avatar []byte
}

type SenderFactory struct {
	cacheSenders map[string]*Sender
}

func (sf *SenderFactory) GetSender(name string) *Sender {
	return sf.cacheSenders[name]
}

func Caller() {
	SenderFactory := SenderFactory{cacheSenders: make(map[string]*Sender)}

	// Create a new sender and add it to the cache
	SenderFactory.cacheSenders["Alice"] = &Sender{
		Name:   "Alice",
		Avatar: make([]byte, 1024*300), // 300kb
	}
	SenderFactory.cacheSenders["Bob"] = &Sender{
		Name:   "Bob",
		Avatar: make([]byte, 1024*400), // 400kb
	}

	fmt.Println([]ChatMessage{
		{
			Content: "Hello",
			Sender:  SenderFactory.GetSender("Alice"),
		},
		{
			Content: "Hi",
			Sender:  SenderFactory.GetSender("Bob"),
		},
		{
			Content: "How are you doing?",
			Sender:  SenderFactory.GetSender("Alice"),
		},
		{
			Content: "I am fine",
			Sender:  SenderFactory.GetSender("Bob"),
		},
	})
	// -> Total memory usage: 300 + 400 = 700kb = 0.7mb
}
