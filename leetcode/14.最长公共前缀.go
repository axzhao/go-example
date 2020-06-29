import "fmt"

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	str := ""
	for i := 0; ; i++ {
		for j := range strs {
			if len(strs[j]) <= i {
				i = -1
				break
			}
			if strs[0][i] != strs[j][i] {
				return str
			}
		}
		if i < 0 {
			break
		}
		str = fmt.Sprintf("%s%s", str, string(strs[0][i]))
	}
	return str
}

// @lc code=end

