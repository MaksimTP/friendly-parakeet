package main

import "fmt"

type Set struct {
	Arr []string
}

func NewSet(arr []string) (*Set, error) {

	uniqueArr := []string{}

	for _, v := range arr {
		isIn := false
		for _, v2 := range uniqueArr {
			if v == v2 {
				isIn = true
			}
		}
		if !isIn {
			uniqueArr = append(uniqueArr, v)
		}
	}

	set := &Set{Arr: uniqueArr}
	return set, nil
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	s, _ := NewSet(arr)
	fmt.Println(s.Arr)
}
