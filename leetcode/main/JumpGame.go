package main

func main() {
	println(jump([]int{2,3,1,1,4}))
}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	n := len(nums)
	longest := 0
	for i := range nums {
		if i <= longest {
			if longest < i+nums[i] {
				longest = i + nums[i]
			}
			if longest >= n-1 {
				return true
			}
		}
	}
	return false
}

func jump(nums []int) int {
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < len(nums) -1; i++ {
		maxPosition = max(maxPosition, nums[i]+i)
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}