package main

import (
	"fmt"
)

/*
1. 分别获取左右边界
2. 获取左边界的时候,如果相等右边界-1
	最终判断左边界有没有大于等于长度
3. 获取右边界的时候,如果相等左边界+1
	最终判断右边界是不是大于等于0

*/

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	leftmost := searchLeft(nums, target)
	rightmost := searchRight(nums, target)
	return []int{leftmost, rightmost}
}

func searchLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		//fmt.Println(left, right, mid)
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println(left)
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func searchRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

func main() {
	nums := []int{2, 2, 3, 3, 3}
	//fmt.Println(searchLeft(nums, 7))
	//fmt.Println(searchRight(nums, 7))
	fmt.Println(searchRange(nums, 3))
}
