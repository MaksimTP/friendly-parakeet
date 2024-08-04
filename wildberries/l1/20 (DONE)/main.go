package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	s := "snow dog sun"
	arr := strings.Split(s, " ")
	slices.Reverse(arr)
	res := strings.Join(arr, " ")
	fmt.Println("Original:", s, "->", "Reversed:", res)
}
