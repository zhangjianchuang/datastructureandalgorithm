package main

func main() {

}
func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func maxSubString(str string) int {
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
