/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]string, len(s))
	top := -1
	for i := range s {
		switch string(s[i]) {
		case "(":
			top++
			stack[top] = string(s[i])
		case "[":
			top++
			stack[top] = string(s[i])
		case "{":
			top++
			stack[top] = string(s[i])
		case ")":
			if top == -1 || stack[top] != "(" {
				return false
			}
			top--
		case "]":
			if top == -1 || stack[top] != "[" {
				return false
			}
			top--
		case "}":
			if top == -1 || stack[top] != "{" {
				return false
			}
			top--
		}
	}
	return top == -1
}
// @lc code=end

