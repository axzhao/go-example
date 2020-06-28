/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	index, length := 0, len(nums)
	for i := 0; i < length-1; i++ {
		if nums[index] == nums[index+1] {
			nums = append(nums[:index], nums[index+1:]...)
			continue
		}
		index++
	}
	return len(nums)
}
// @lc code=end

