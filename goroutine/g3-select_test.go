package goroutine

import (
	"fmt"
	"time"
)

func ExampleSelect1() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	ch1 = nil // disables this channel
	select {
	case <-ch1:
		fmt.Println("Received from ch1") // will not happen
	case <-ch2:
		fmt.Println("Received from ch2")
	default:
		fmt.Println("Nothing available")
	}

	// Output:
}

func ExampleSelect2() {
	AFP := make(chan string)
	select {
	case news := <-AFP:
		fmt.Println(news)
	case <-time.After(time.Minute):
		fmt.Println("Time out: No news in one minute")
	}

	// Output:
}
