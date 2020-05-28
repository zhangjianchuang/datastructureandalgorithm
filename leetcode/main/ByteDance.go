package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(maxSubStr(" "))
	//fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}

func maxSubStr(str string) int {
	// 结果，也就是不重复的串最大长度
	ans := 0
	// 滑动的窗口，代表128个英文字符，如果有中文字符就不够用了
	index := make([]int, 128)
	// 上一次出现重复的位置
	i := 0
	// 字符串转换成的 byte 数组
	s := []byte(str)
	for j := 0; j < len(str); j++ {
		// 先在数组中寻找上一次这个字符出现的位置，如果已经出现过，
		i = max(index[s[j]], i)
		// 最大值就是 字符串当前遍历到的位置 - 上次重复出现的位置 + 1
		// 意思就是把出现重复的字符以及之前的字符全部抛弃，剩下的长度
		ans = max(ans, j-i+1)
		// 记录当前字符出现的位置
		index[s[j]] = j + 1
	}
	return ans
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func longestCommonPrefix(strs []string) string {
	result := ""
	if len(strs) == 0 {
		return result
	}
	if strs[0] == "" {
		return result
	}
	for j := 1; j <= len(strs[0]); j++ {
		result = strs[0][0:j]
		for i := range strs {
			if !strings.HasPrefix(strs[i], result) {
				return result[0 : j-1]
			}
		}
	}
	return result
}
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 || len(s1) > len(s2) {
		return false
	}
	window1 := make([]int, 26)
	window2 := make([]int, 26)
	for i := range s1 {
		println(s1[i])
		println(s1[i] - 'a')
		window1[s1[i]-'a']++
		window2[s2[i]-'a']++
	}
	diff := make([]int, 26)
	for i := range diff {
		diff[i] = window2[i] - window1[i]
	}
	for i := len(s1); i < len(s2); i++ {
		if same(diff) {
			return true
		}
		diff[s2[i-len(s1)]-'a']--
		diff[s2[i]-'a']++
	}
	return same(diff)
}
func same(diff []int) bool {
	for i := range diff {
		if diff[i] != 0 {
			return false
		}
	}
	return true
}

func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return ""
	}
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	muls := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			mul := int((num1[i] - '0') * (num2[j] - '0'))
			mul += muls[i+j+1]
			muls[i+j+1] = mul % 10
			muls[i+j] += mul / 10
		}
	}
	start := 0
	for start < len(muls) && muls[start] == 0 {
		start++
	}
	var r strings.Builder
	for start < len(muls) {
		r.WriteString(strconv.Itoa(muls[start]))
		start++
	}
	return r.String()
}
