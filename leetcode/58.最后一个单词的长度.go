/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	lastCount, count := 0, 0
	for i := range s {
		if string(s[i]) == " " {
			count = 0
			continue
		}
		count++
		lastCount = count
	}
	return lastCount
}
// @lc code=end

