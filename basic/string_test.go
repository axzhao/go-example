package basic

import (
	"fmt"
	"unicode/utf8"
)

func ExampleString() {
	s := "世界abc"
	fmt.Println([]byte(s))
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	// Output:
}
