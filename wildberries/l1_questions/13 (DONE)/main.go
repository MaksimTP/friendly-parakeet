package main

import (
	"fmt"
)

func someAction(v []int8, b int8) {
	fmt.Printf("%p\n", v)
	v[0] = 100
	v = append(v, b)
	fmt.Printf("%p\n", v)
}
func main() {
	var a = []int8{1, 2, 3, 4, 5}
	fmt.Printf("%p, %d, %d\n", a, len(a), cap(a))
	someAction(a, 6)
	fmt.Println(a)
}
