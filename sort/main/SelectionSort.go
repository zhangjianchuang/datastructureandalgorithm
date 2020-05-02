package main

import "fmt"

func selectionSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		arr[minIdx], arr[i] = arr[i], arr[minIdx]
	}
}

func main() {
	arr := []int{4, 5, 6, 1, 2, 3}
	selectionSort(arr)
	fmt.Println(arr)
}
