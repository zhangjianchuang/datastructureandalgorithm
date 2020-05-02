package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	if n <= 0 {
		return
	}
	for i := 1; i < n; i++ {
		val := arr[i]
		j := i - 1
		// 把所有比当前值大的都向后挪一位
		for ; j >= 0 && val < arr[j]; j-- {
			arr[j+1] = arr[j]
		}
		// 插入操作
		arr[j+1] = val
	}
}

func main() {
	arr := []int{4, 5, 6, 4, 1, 2, 3}
	insertionSort(arr)
	fmt.Println(arr)
}
