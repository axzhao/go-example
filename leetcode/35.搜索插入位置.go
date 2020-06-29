/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	low, mid, high := 0, (len(nums)-1)/2, len(nums)-1
	if target > nums[high] {
		return len(nums)
	}
	if target < nums[low] {
		return 0
	}
	for high >= low {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			low = mid + 1
		}
		if nums[mid] > target {
			high = mid - 1
		}
		mid = (high + low) / 2
	}
	return mid + 1 
}
// @lc code=end

