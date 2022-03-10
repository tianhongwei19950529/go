package main

func lengthOfLongestSubstring(s string) int {
	/*
		1. 定义左右指针,右指针从-1开始(从0开始不可以)
		2. 循环这个字符串
		3.
	*/

	left, right := 0, -1
	slen := len(s)
	m := make(map[byte]int)
	for i := 0; i < slen; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for right < slen && m[s[right+1]] == 0 {
			
		}
	}
	return 0
}
