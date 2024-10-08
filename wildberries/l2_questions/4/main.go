package main

import "fmt"

func main() {
	ch := make(chan int, 0)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
