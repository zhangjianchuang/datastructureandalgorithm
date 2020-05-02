package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	begin := time.Now().UnixNano()
	fmt.Println(numberOf2s(559366752))
	fmt.Printf("我的：%d\n",time.Now().UnixNano() - begin)

	begin = time.Now().UnixNano()
	fmt.Println(numberOf2sInRange(559366752))
	fmt.Printf("别人的：%d",time.Now().UnixNano() - begin)
}

func numberOf2sInRange(n int) int {
	if n == 0 {
		return 0
	}
	digit := int(math.Log10(float64(n))) + 1
	dp := new([10][2]int)
	if n%10 >= 2 {
		dp[1][0] = 1
	} else {
		dp[1][0] = 0
	}
	dp[1][1] = 1
	for i := 2; i <= digit; i++ {
		k := n/int(math.Pow10(i-1)) % 10
		dp[i][0] = k*dp[i-1][1] + dp[i-1][0]
		if k == 2 {
			dp[i][0] += n%int(math.Pow10(i-1)) + 1
		} else if k > 2 {
			dp[i][0] += int(math.Pow10(i - 1))
		}
		dp[i][1] = 10*dp[i-1][1] + int(math.Pow10(i-1))
	}
	return dp[digit][0]
}

func numberOf2s(n int) int {
	count := 0
	for i := 0; i <= n; i++ {
		num := i
		for num > 0 {
			if num%10 == 2 {
				count++
			}
			num /= 10
		}
	}
	return count
}
