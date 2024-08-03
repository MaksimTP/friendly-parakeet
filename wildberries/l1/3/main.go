package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	var res *int = new(int)
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(val int, wg *sync.WaitGroup) {
			defer wg.Done()
			*res += val * val
		}(arr[i], &wg)
	}
	wg.Wait()
	fmt.Println(*res)

}
