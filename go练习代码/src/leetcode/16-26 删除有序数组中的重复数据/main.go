package main

/*
	1. 双指针(快慢指针)的另一种实现方式.
	2. 由快指针遍历数组,判断当前值是不是和前一个值相等.
	3. 如果不相等的话将快指针的值付给慢指针.
	4. 计数器+1
*/
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {
}
