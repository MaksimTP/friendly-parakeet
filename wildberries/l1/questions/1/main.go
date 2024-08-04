package main

import (
	"bytes"
)

func main() {

	// #1. Using a buffer
	strings := []string{"I", "hate", "this", "world"}
	buffer := bytes.Buffer{}
	for _, v := range strings {
		buffer.WriteString(v)
	}
	// fmt.Println(buffer.String())

	// #2. Using a copy function

	bs := make([]byte, 100)
	bl := 0

	for _, v := range strings {
		bl += copy(bs[bl:], []byte(v))
	}

	// fmt.Println(string(bs[:]))

	// #3. Just concat

	res := ""

	for _, v := range strings {
		res += v
	}
}
