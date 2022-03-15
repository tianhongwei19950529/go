package main

import "fmt"

func longestPalindrome(s string) string {
	start, end := 0, 0
	n := len(s)
	for i := 0; i < n; i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		//仅仅是因为长度为2???
		left2, right2 := expandAroundCenter(s, i, i+1)
		if (right1 - left1) > (end - start) {
			start, end = left1, right1
		}
		if (right2 - left2) > (end - start) {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left -= 1
		right += 1
	}
	return left + 1, right - 1
}

func main() {
	s := "b"
	fmt.Println(longestPalindrome(s))
}
