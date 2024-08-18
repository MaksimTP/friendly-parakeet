package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortStr(str string) string {
	s := []rune(strings.Clone(str))

	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })

	return string(s)
}

func elemInSlice(arr []string, s string) (string, bool) {
	for _, v := range arr {
		if sortStr(v) == sortStr(s) {
			return v, true
		}
	}
	return "", false
}

func Anagrams(arr []string) map[string][]string {
	m := make(map[string][]string)

	firstOccurences := make([]string, 0)

	for _, v := range arr {
		str, seen := elemInSlice(firstOccurences, v)
		if !seen {
			firstOccurences = append(firstOccurences, v)
			m[v] = make([]string, 0)
			m[v] = append(m[v], v)
		} else {
			m[str] = append(m[str], v)
		}
	}

	for k, _ := range m {
		sort.Strings(m[k])
	}

	return m
}

func main() {
	s := []string{"пятак", "пятка", "тяпка", "слиток", "листок", "столик"}
	m := Anagrams(s)
	fmt.Println(m)
}
