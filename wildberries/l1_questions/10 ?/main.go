package main

import "fmt"

func update(p *int) {
	b := 2
	p = &b
	fmt.Println(*p, p)
}
func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p, p)
	update(p)
	fmt.Println(*p, p)
}
