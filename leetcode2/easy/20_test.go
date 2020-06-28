package easy

import "fmt"

func Example20() {

	fmt.Println(isValid("(("))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))

	// Output:
}

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
