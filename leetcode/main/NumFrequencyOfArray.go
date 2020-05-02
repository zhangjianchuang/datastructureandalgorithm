package main

func main() {
	numbers := singleNumbers([]int{4, 1, 4, 6})
	println(numbers[0])
	println(numbers[1])
}
func singleNumbers(nums []int) []int {
	xorSum := 0
	ret := []int{0, 0}
	for x := range nums {
		xorSum ^= nums[x]
	}
	lowbit := xorSum & (-xorSum)
	for x := range nums {
		if nums[x]&lowbit > 0 {
			ret[0] ^= nums[x]
		} else {
			ret[1] ^= nums[x]
		}
	}
	return ret
}
