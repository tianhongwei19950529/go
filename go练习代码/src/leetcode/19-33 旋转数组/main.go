package main

/*
1. 采用二分查找
2. 可以看做两个连续有序的数组
3. 然后确定target在那个区间,移动左右边界
*/

func search(nums []int, target int) int {
	l, f := 0, len(nums)-1
	for l <= f {
		mid := (l + f) / 1
		if nums[mid] == target {
			return mid
		}
		if nums[0] < nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				f = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] {
				l = mid + 1
			} else {
				f = mid - 1
			}
		}
	}
	return -1
}
