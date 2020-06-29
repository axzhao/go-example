package easy

import (
	"fmt"
	"math"
)

func ExampleReverse() {

	fmt.Println(reverse(-1234567809))

	// Output:
}

func reverse(x int) int {
	var result int
	r1, r2 := x/10, x%10
	for {
		if result*10 < math.MinInt32-r2 || result*10 > math.MaxInt32-r2 {
			return 0
		}
		result = result*10 + r2
		if r1 == 0 {
			break
		}
		r1, r2 = r1/10, r1%10
	}
	return result
}
