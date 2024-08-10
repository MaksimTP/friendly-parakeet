package main

import "fmt"

type Element interface {
	accept(Visitor)
}

type Visitor interface {
	visit(Element)
}

type ml int

type Drink struct {
	Volume ml
}

type Water struct{ Drink }

func (e Water) accept(v Visitor) {
	v.visit(e)
}

type Cola struct{ Drink }

func (e Cola) accept(v Visitor) {
	v.visit(e)
}

type Pepsi struct{ Drink }

func (e Pepsi) accept(v Visitor) {
	v.visit(e)
}

type Watcher struct{}

func (watch Watcher) visit(e Element) {
	fmt.Printf("%+v\n", e)
}

func main() {
	elems := []Element{Water{Drink{500}}, Cola{Drink{1500}}, Pepsi{Drink{300}}}

	w := Watcher{}

	for _, v := range elems {
		w.visit(v)
	}
}
