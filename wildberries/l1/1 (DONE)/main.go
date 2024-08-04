package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры
// Human (аналог наследования).

type Human struct {
	name string
	age  int
}

type Action struct {
	Human
	job string
}

func (h *Human) changeName(name string) {
	h.name = name
}

func (h *Human) changeAge(age int) {
	if age >= 0 {
		h.age = age
	}
}

func main() {
	worker := &Action{
		Human: Human{name: "Tom", age: 20},
		job:   "Actor",
	}
	worker.changeAge(30)
	fmt.Println(worker.age, worker.job)
}
