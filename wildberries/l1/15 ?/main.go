package main

import (
	"fmt"
)

var justString string

func createHugeString(num int) (res string) {

	for i := 0; i < num; i++ {
		res += "h"
	}
	fmt.Println(len(res))
	return
}

func someFunc() {
	v := createHugeString(1 << 13)
	fmt.Println(len(v))
	justString = v[:100]
}

func main() {
	someFunc()
	fmt.Println(justString, len(justString))
}
