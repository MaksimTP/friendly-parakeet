package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[0] = 1
	m[1] = 124
	m[2] = 281
	for k, v := range m {
		fmt.Println(k, v)
	}
}
