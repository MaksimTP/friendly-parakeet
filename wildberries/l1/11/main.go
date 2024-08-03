package main

import "fmt"

type Set struct {
	Arr []interface{}
}

func (s Set) isIn(value interface{}) bool {
	for _, v := range s.Arr {
		if v == value {
			return true
		}
	}
	return false
}

func (s Set) Intersection(s2 Set) Set {
	newS := Set{}
	for _, v := range s.Arr {
		if !newS.isIn(v) {
			newS.Arr = append(newS.Arr, v)
		}
	}
	for _, v := range s2.Arr {
		if !newS.isIn(v) {
			newS.Arr = append(newS.Arr, v)
		}
	}

	return newS
}

func main() {
	s1 := Set{}
	s1.Arr = []interface{}{1, 2, 3, 4, 5, 6}

	s2 := Set{}
	s2.Arr = []interface{}{2, 3, 4, 5, 6, 7, 3213, 4, 5}

	s3 := s1.Intersection(s2)
	fmt.Println(s3)
}
