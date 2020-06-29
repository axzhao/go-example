package easy

import (
	"fmt"
	"strconv"
)

func Example38() {

	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))

	// Output:
}

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
