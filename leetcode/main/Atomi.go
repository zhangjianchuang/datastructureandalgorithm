package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(int('1' - 48))
}

func myAtoi(str string) int {
	var maxNumber int64 = math.MaxInt32
	str = trim(str)
	flag := false
	var num int64 = 0
	if len(str) == 0 {
		return int(num)
	}

	if str[0] == '-' {
		flag = true
		str = str[1:]
	} else if str[0] == '+' {
		flag = false
		str = str[1:]
	}

	for i := 0; i < len(str) && str[i] >= '0' && str[i] <= '9'; i++ {
		num = num*10 + int64(str[i]-48)
		if num > maxNumber {
			if flag {
				num = maxNumber + 1
				return int(-num)
			} else {
				num = maxNumber
				return int(num)
			}
		}
	}
	if flag {
		return int(-num)
	}
	return int(num)
}

func trim(str string) string {
	for i := range str {
		if str[i] != ' ' {
			return str[i:]
		}
	}
	return str
}
