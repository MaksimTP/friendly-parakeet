package main

import "fmt"

// import "sort"

func groupValues(arr []float32, step int) map[int][]float32 {
	m := make(map[int][]float32)

	for _, v := range arr {
		key := int(v) / step * step
		if key < 0 {
			key -= step
		}
		m[key] = append(m[key], v)
	}
	fmt.Println(m)
	return m
}

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groupValues(arr, 10)
}
