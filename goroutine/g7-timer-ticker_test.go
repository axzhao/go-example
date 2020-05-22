package goroutine

import (
	"fmt"
	"log"
	"time"
)

func ExampleTimeoutTimer() {
	AFP := make(chan struct{})
	select {
	case news := <-AFP:
		fmt.Println(news)
	case <-time.After(time.Hour):
		fmt.Println("No news in an hour.")
	}

	for alive := true; alive; {
		timer := time.NewTimer(time.Hour)
		select {
		case news := <-AFP:
			timer.Stop()
			fmt.Println(news)
		case <-timer.C:
			alive = false
			fmt.Println("No news in an hour. Service aborting.")
		}
	}

	// Output:
}

func ExampleRepeatTicker() {
	// The underlying time.
	// Ticker will not be recovered by the garbage collector.
	// If this is a concern, use time.NewTicker instead and call its Stop method when the ticker is no longer needed.
	go func() {
		for now := range time.Tick(time.Minute) {
			fmt.Println(now, statusUpdate())
		}
	}()
	// Output:
}

func ExampleAfterFunc() {
	timer = time.AfterFunc(time.Minute, func() {
		log.Println("Foo run for more than a minute.")
	})
	defer timer.Stop()

	// Do heavy work
	// Output:
}
