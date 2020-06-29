package easy

import (
	"fmt"
)

func ExampleIsPalindrome() {
	for _, r := range []struct {
		actual   bool
		expected bool
	}{
		{isPalindrome(-121), false},
		{isPalindrome(121), true},
		{isPalindrome(1001), true},
		{isPalindrome(10), false},
	} {
		if r.actual != r.expected {
			fmt.Println("err")
		}
	}

	for _, r := range []struct {
		actual   bool
		expected bool
	}{
		{isPalindromeStr("-121"), false},
		{isPalindromeStr("121"), true},
		{isPalindromeStr("1001"), true},
		{isPalindromeStr("10"), false},
	} {
		if r.actual != r.expected {
			fmt.Println("err")
		}
	}
	// Output:
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	var result int
	r1, r2 := x/10, x%10
	for {
		result = result*10 + r2
		if r1 == 0 {
			break
		}
		r1, r2 = r1/10, r1%10
	}
	return result == x
}

func isPalindromeStr(str string) bool {
	if len(str) <= 1 {
		return true
	}
	for i := 0; i <= len(str)/2; i++ {
		if str[i] == str[len(str)-1-i] {
			continue
		}
		return false
	}
	return true
}
