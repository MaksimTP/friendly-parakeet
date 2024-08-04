package main

import "fmt"

func binarySearch(arr []int, val int) (int, error) {
	low := 0
	high := len(arr)

	iterCnt := 0

	for low <= high {
		searchIndex := (high + low) / 2
		if arr[searchIndex] == val {
			return searchIndex, nil
		} else if val < arr[searchIndex] {
			high = searchIndex
		} else if val > arr[searchIndex] {
			low = searchIndex
		}
		iterCnt++
		if iterCnt > len(arr) {
			break
		}

	}
	return -1, fmt.Errorf("cant find")
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wantedValue := 1123
	ind, err := binarySearch(arr, wantedValue)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(arr[ind], wantedValue)
	}

}
