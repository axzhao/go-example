package basic

import (
	"fmt"
)

func ExampleSwitch() {

	s := "("

	for i := range s {
		switch string(s[i]) {
		case "(":
			fmt.Println("666")
			fallthrough
		case "[":
			fmt.Println("555")
			fallthrough
		case "{":
			fmt.Println("111")
		case ")":
			fmt.Println("222")
		case "]":
			fmt.Println("333")
		case "}":
			fmt.Println("444")
		}
	}

	// Output:
}
