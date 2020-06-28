/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" || needle == haystack {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	index := -1
	for i := 0; i < len(haystack); i++ {
		index = i
		for j := 0; j < len(needle); j++ {
			if i+j >= len(haystack) {
				return -1
			}
			if string(haystack[i+j]) == string(needle[j]) {
				continue
			}
			index = -1
			break
		}
		if index != -1 {
			return index
		}
	}
	return -1
}
// @lc code=end

