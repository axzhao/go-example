import "math"

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	var result int
	r1, r2 := x/10, x%10
	for {
		if result*10 < math.MinInt32-r2 || result*10 > math.MaxInt32-r2 {
			return 0
		}
		result = result*10 + r2
		if r1 == 0 {
			break
		}
		r1, r2 = r1/10, r1%10
	}
	return result
}

// @lc code=end

