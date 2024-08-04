package main

import "fmt"

func reverse(s string) string {
	r := []rune(s)
	res := ""
	for _, v := range r {
		res = string(v) + res
	}
	return res
}

func main() {
	x := "hell😊o"
	fmt.Println(x, reverse(x))
}
