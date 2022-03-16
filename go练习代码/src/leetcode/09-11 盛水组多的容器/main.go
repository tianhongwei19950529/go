package main

import "fmt"

/*
思路:  双指针,两边往中间走,底是两个指针的距离 高是两个指针对应的短的长度
*/

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxnum := 0
	for left < right {
		hei := 0
		if height[left] <= height[right] {
			hei = height[left]
		} else {
			hei = height[right]
		}
		maxnum = max(maxnum, hei*(right-left))
		if height[left] <= height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return maxnum
}

func min(left, right int) int {
	if left > right {
		return right
	}
	return left
}
func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}

func main() {
	numlist := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(numlist))
}
