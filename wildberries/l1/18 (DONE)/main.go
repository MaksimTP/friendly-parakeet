package main

import (
	"fmt"
	"sync"
)

type ConcurrentCounter struct {
	Count int
	mu    sync.Mutex
}

func (c *ConcurrentCounter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Count++
}

func (c *ConcurrentCounter) display() {
	fmt.Println(c.Count)
}

func main() {
	counter := ConcurrentCounter{}
	var wg sync.WaitGroup
	for i := 0; i < 90000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.increment()
		}()
	}

	wg.Wait()
	counter.display()
}
