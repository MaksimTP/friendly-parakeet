package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(val int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(val * val)
		}(arr[i], &wg)
	}

	wg.Wait()
}
