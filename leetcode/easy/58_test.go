package easy

import "fmt"

func Example58() {
	fmt.Println(lengthOfLastWord(("")))

	// Output:
}
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
