/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	var result int
	r1, r2 := x/10, x%10
	for {
		result = result*10 + r2
		if r1 == 0 {
			break
		}
		r1, r2 = r1/10, r1%10
	}
	return result == x
}

// @lc code=end

