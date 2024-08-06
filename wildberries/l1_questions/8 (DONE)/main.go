package main

import "fmt"

func main() {
	a := new(map[int]int)
	b := make(map[int]int)
	fmt.Printf("%T %T\n", a, b)
}
