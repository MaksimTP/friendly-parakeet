package main

import "fmt"

func swapValues(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {

	// #1
	a := 2
	b := 3
	fmt.Println(a, b)
	swapValues(&a, &b)
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
