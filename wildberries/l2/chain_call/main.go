package main

import "fmt"

type Handler interface {
	HandleMessage()
}

type Handler1 struct {
	successor Handler
}

func (h Handler1) HandleMessage() {
	if h.successor == nil {
		h.HandleMethod()
	} else {
		h.successor.HandleMessage()
	}
}

func (h Handler1) HandleMethod() {
	fmt.Println("Handler1 doing something")
}

type Handler2 struct {
	successor Handler
}

func (h Handler2) HandleMessage() {
	if h.successor == nil {
		h.HandleMethod()
	} else {
		h.successor.HandleMessage()
	}
}

func (h Handler2) HandleMethod() {
	fmt.Println("Handler2 doing something")
}

func main() {
	h1 := Handler1{}
	h2 := Handler2{}
	h1.successor = h2
	h1.HandleMessage()
}
