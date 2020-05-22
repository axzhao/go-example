package goroutine

import "fmt"

// go test -race [packages]
// go run -race [packages]
func ExampleRace() {
	// wait := make(chan struct{})
	n := 0
	go func() {
		n++ // read, increment, write
		// close(wait)
	}()
	n++ // conflicting access
	// <-wait
	fmt.Println(n) // Output: <unspecified>
	// Output:
}

func ExampleRace1() {
	// 1. it passes the data from one goroutine to another,
	// 2. and it acts as a point of synchronization.
	ch := make(chan int)
	go func() {
		n := 0 // A local variable is only visible to one goroutine.
		n++
		ch <- n // The data leaves one goroutine...
	}()
	n := <-ch // ...and arrives safely in another.
	n++
	fmt.Println(n) // Output: 2
	// Output:
}
