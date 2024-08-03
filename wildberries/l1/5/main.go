package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	ch := make(chan int, 5)

	var N int = 1
	timeout := time.After(time.Duration(N) * time.Second)

	go func() {
		for {
			select {
			case x := <-ch:
				fmt.Println(x)
			case <-timeout:
				fmt.Println("timeout happened, exiting...")
				os.Exit(0)
			}
		}
	}()

	for {
		ch <- rand.Int()
	}

}
