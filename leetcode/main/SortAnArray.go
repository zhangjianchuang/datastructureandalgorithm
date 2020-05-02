package main

import "fmt"

func main() {
	array := sortArray([]int{5, 2, 3, 1})
	fmt.Println(array)
}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(arr []int, begin, end int) {
	if begin >= end {
		return
	}
	par := getPar(arr, begin, end)
	quickSort(arr, begin, par-1)
	quickSort(arr, par+1, end)
}

func getPar(arr []int, begin int, end int) int {
	base := arr[begin]
	for begin < end {
		for arr[end] >= base && end > begin {
			end--
		}
		arr[begin] = arr[end]
		for arr[begin] <= base && end > begin {
			begin++
		}
		arr[end] = arr[begin]
	}
	arr[end] = base
	return end
}
