package easy

import "fmt"

func Example14() {
	fmt.Println(longestCommonPrefix([]string{}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))

	// Output:
	// 223
}

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
