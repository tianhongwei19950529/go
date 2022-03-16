package main

import "fmt"

//暴力方法
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	n := len(needle)
	n2 := len(haystack)
	for i := 0; i < n2; i++ {
		if haystack[i] == needle[0] {
			if i+n > n2 {
				return -1
			} else if needle == haystack[i:i+n] {
				return i
			}
		}
	}
	return -1
}

//kmp????
func strStr1(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	pi := make([]int, m)
	//[0,0,0,0]
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			pi[j] = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != haystack[j] {
			pi[j] = pi[j-1]
		}
		if haystack[i] == haystack[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
func main() {
	haystack, needle := "a", "a"
	fmt.Println(strStr(haystack, needle))
}
