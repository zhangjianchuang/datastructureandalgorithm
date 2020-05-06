package main

func main() {

}

func mincostTickets(days []int, costs []int) int {
	memo := [366]int{}
	dayM := map[int]bool{}
	for _, d := range days {
		dayM[d] = true
	}

	var dp func(day int) int
	dp = func(day int) int {
		if day > 365 {
			return 0
		}
		if memo[day] > 0 {
			return memo[day]
		}
		if dayM[day] {
			memo[day] = min(min(dp(day+1)+costs[0], dp(day+7)+costs[1]), dp(day+30)+costs[2])
		} else {
			memo[day] = dp(day + 1)
		}
		return memo[day]
	}
	return dp(1)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
