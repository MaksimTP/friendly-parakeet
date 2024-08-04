package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

func main() {
	var l sync.Mutex

	m := make(map[string]int, 5)
	var wg sync.WaitGroup
	arrKeys := []string{"a", "b", "c", "d", "e"}
	arrValues := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(arrKeys); i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			l.Lock()
			defer wg.Done()
			defer l.Unlock()
			m[arrKeys[i]] = arrValues[i]
		}(&wg)
	}
	wg.Wait()
	fmt.Println(m)
}
