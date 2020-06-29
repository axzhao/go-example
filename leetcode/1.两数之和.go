/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		panic("")
	}
	tmp := make(map[int]int, len(nums))
	for i, v := range nums {
		if vv, ok := tmp[v]; ok {
			return []int{vv, i}
		}
		tmp[target-v] = i
	}
	return []int{}
}

// @lc code=end
