package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for {
			res := <-in
			fmt.Printf("source: %d ", res)
			out <- res * res
		}
	}()

	go func() {
		for {
			in <- rand.Int() % 15
			time.Sleep(time.Second)
		}
	}()

	for {
		fmt.Printf("squared: %d\n", <-out)
	}

}
