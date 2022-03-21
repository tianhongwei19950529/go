package main

import (
	"fmt"
)

/*
1. 利用动归的思想
2. max取第一个，然后从第一个开始取数，max（nums[i-1]+nums[i],num[i]）,也就是说先比较当前加上前面的有没有比当前的大
*/
func maxSubArray(nums []int) int {
	max := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func maxSubArray1(nums []int) int {
	n := len(nums)
	targes := make([]int, n)
	targes[0] = nums[0]
	for i := 1; i < n; i++ {
		gold := max(nums[i]+targes[i-1], nums[i])
		targes[i] = gold
	}
	fmt.Println(targes)
	return 0
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSubArray1(nums)
}
