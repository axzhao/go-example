/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	if strconv.Itoa(n) == "1" {
		return "1"
	}
	str := countAndSay(n - 1)
	count := 0
	flag, result := string(str[0]), ""
	for i := range str {
		if flag != string(str[i]) {
			result += fmt.Sprintf("%d%s", count, flag)
			flag = string(str[i])
			count = 1
		} else {
			count++
		}
	}
	result += fmt.Sprintf("%d%s", count, flag)
	return result
}
// @lc code=end

