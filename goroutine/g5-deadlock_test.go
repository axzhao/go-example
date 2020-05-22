package goroutine

import "fmt"

func ExampleDeadlock() {
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
	// Output:
}
