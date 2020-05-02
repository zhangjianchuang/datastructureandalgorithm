package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(maxSubStr("asdfa"))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(checkInclusion("123", "1323"))
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
	//state := make([]bool,len(s1))
	//for i := range s2 {
	//	if s2[i] == s1[j] {
	//		if j == len(s1)-1 {
	//			return true
	//		} else {
	//			j++
	//		}
	//	} else {
	//		if s2[i] == s1[0] {
	//			j = 1
	//		} else {
	//			j = 0
	//		}
	//	}
	//}
	return false
}
