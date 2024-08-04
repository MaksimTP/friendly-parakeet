package main

import "fmt"

func isUnique(s string) bool {
	seen := make(map[rune]bool)

	for _, v := range s {
		if value, found := seen[v]; !found {
			seen[v] = value
		} else {
			return false
		}
	}
	return true
}

func main() {
	a1 := "abcd"
	a2 := "aAs"
	a3 := "aas"
	fmt.Println(isUnique(a1))
	fmt.Println(isUnique(a2))
	fmt.Println(isUnique(a3))
}
