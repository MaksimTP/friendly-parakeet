package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	c := make(chan interface{}, len(channels))

	for _, channel := range channels {
		val := <-channel
		c <- val
	}
	return c
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
	)

	fmt.Printf("Done after %v\n", time.Since(start))
}
