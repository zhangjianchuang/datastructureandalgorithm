package main

import (
	"fmt"
)

func main() {
	array := []int{3, 6, 1, 4, 2, 8}
	fmt.Println(array)
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}

func quickSort(array []int, left, right int) {
	if left >= right {
		return
	}
	index := partition(array, left, right)
	quickSort(array, left, index-1)
	quickSort(array, index+1, right)
}

func partition(array []int, left, right int) int {
	baseNum := array[left]
	for left < right {
		for array[right] >= baseNum && right > left {
			right--
		}
		array[left] = array[right]
		for array[left] <= baseNum && right > left {
			left++
		}
		array[right] = array[left]
	}
	array[right] = baseNum
	return right
}