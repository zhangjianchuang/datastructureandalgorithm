package main

import (
	"fmt"
)

// mergerSort
func mergerSort(arr []int, a, b int) {
	if b-a <= 1 {
		return
	}

	c := (a + b) / 2
	// 先分开，分成一个一个的元素
	mergerSort(arr, a, c)
	mergerSort(arr, c, b)

	// 归并步骤
	arrLeft := make([]int, c-a)
	arrRight := make([]int, b-c)
	copy(arrLeft, arr[a:c])
	copy(arrRight, arr[c:b])
	// 左边进行
	i := 0
	// 右边进行
	j := 0
	// c 是中间分开的下标
	// a 是起始位置
	// b 是结束位置
	// arr 是新的合并数组
	for k := a; k < b; k++ {
		if i >= c-a {
			arr[k] = arrRight[j]
			j++
		} else if j >= b-c {
			arr[k] = arrLeft[i]
			i++
		} else if arrLeft[i] < arrRight[j] {
			arr[k] = arrLeft[i]
			i++
		} else {
			arr[k] = arrRight[j]
			j++
		}
	}
}

func main() {
	// 测试代码
	arr := []int{9, 8, 7, 6, 5, 1, 2, 3, 4, 0}
	fmt.Println(arr)
	mergerSort(arr, 0, len(arr))
	fmt.Println(arr)
}