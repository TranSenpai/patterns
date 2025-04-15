package chainofresponsibility

import "fmt"

// What
//  - is a behavioral design pattern that lets you pass requests along a chain of handlers.
//  - Upon receiving (khi nhận) a request, each handler decides either to process the request or to
//  - pass it to the next handler in the chain.

//                           handler1 (middleware)
// Example: Request login -> Authen(JWT, OAuth...) -> Ordering System

// When
//  - difficult request requiring multiple steps to process
//  - reuse some steps of process

// Why
//  - to separate the sender and receiver in the request processing
//  - allow multiple objects to have the opportunity to process the request
//    without knowing in advance who will handle it
//  - simplify the handle logic by breaking it into smaller pieces
//    each handler is responsible for a single task

// Problem:
type WebCrawler1 struct{}

func (WebCrawler1) Crawler(url string) {
	fmt.Println("A simple process of crawling a url:", url)
	fmt.Println("Check the url")
	fmt.Println("Fetch url content")
	fmt.Println("Extract information from content")
	fmt.Println("Save information to database")
}

// Solution:
// How:

type Context struct {
	url     string
	content string
	data    any
}

type Handler func(*Context) error

func CheckingUrlhandler(ctx *Context) error {
	fmt.Printf("Check the url: %s\n", ctx.url)
	return nil
}

func FetchContentHandler(ctx *Context) error {
	fmt.Printf("Fetch url content: %s\n", ctx.url)
	ctx.content = "some content from url"
	return nil
}

func ExtractInformationHandler(ctx *Context) error {
	fmt.Printf("Extract information from content: %s\n", ctx.content)
	ctx.data = map[string]interface{}{"title": "Apple", "price": 10.0}
	return nil
}

func SaveInformationHandler(ctx *Context) error {
	fmt.Printf("Save information to database: %s\n", ctx.data)
	return nil
}

type HandlerNode struct {
	hdl  Handler
	next *HandlerNode
}

func (h *HandlerNode) Handle(url string) error {
	ctx := Context{url: url}

	if h == nil || h.hdl == nil {
		return nil
	}

	// gọi handder thực thi việc gì đó, check nếu có lỗi thì trả về lỗi
	// không thì đệ quy gọi tiếp đến khi next == nil
	if err := h.hdl(&ctx); err != nil {
		return err
	}
	return h.next.Handle(url)
}

func NewCrawler(handlers ...Handler) HandlerNode {
	node := HandlerNode{}

	if len(handlers) > 0 {
		node.hdl = handlers[0]
	}

	currentNode := &node

	for i := 1; i < len(handlers); i++ {
		currentNode.next = &HandlerNode{hdl: handlers[i]}
		currentNode = currentNode.next
	}

	return node
}

type WebCrawler struct {
	handler HandlerNode
}

func (wc WebCrawler) Crawl(url string) {
	if err := wc.handler.Handle(url); err != nil {
		fmt.Println(err)
	}
}

func Caller() {
	WebCrawler{
		handler: NewCrawler(
			CheckingUrlhandler,
			FetchContentHandler,
			ExtractInformationHandler,
			SaveInformationHandler,
		),
	}.Crawl("http://example.com")

}
