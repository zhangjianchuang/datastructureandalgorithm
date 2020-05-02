package main

import "fmt"

func bubbleSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			fmt.Println("没有要排序的了，直接返回了")
			return
		}
	}
}

func main() {
	arr := []int{3, 4, 5, 1, 2}
	bubbleSort(arr)
	fmt.Println(arr)
	for i := range arr {
		fmt.Println(i)
	}
}
