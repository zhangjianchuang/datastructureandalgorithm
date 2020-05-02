package main

func main() {
	println(numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
}

// 遍历数组，先取k长度看是否符合，然后加长看是否符合
func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	odd := make([]int, n+2)
	ans := 0
	cnt := 1
	for i := range nums {
		if nums[i]%2 == 1 {
			odd[cnt] = i
			cnt++
		}
	}
	odd[0] = -1
	odd[cnt] = n
	for i := 1; i+k <= cnt; i++ {
		ans += (odd[i] - odd[i-1]) * (odd[i+k] - odd[i+k-1])
	}
	return ans
}
