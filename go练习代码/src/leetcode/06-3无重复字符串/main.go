package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	/*
		1. 需要一个哈希集合 记录字符串是不是出现过
		2. 需要左右指针,左指针是for循环中的i代表,右指针必须有变量代表.
		3. 不需要每次都重新移动右指针
		4. 每一次都要记录最大值
	*/
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	right, count := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for right+1 < n && m[s[right+1]] == 0 {
			m[s[right+1]]++
			right++
		}
		count = max(count, right-i+1)
	}
	return count
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
