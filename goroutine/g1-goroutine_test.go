package goroutine

import (
	"fmt"
	"time"
)

func ExampleGoroutine1() {

	go fmt.Println("Hello from another goroutine")
	fmt.Println("Hello from main goroutine")

	// At this point the program execution stops and all
	// active goroutines are killed.
	// time.Sleep(time.Second) // give the other goroutine time to finish

	// Output:
}

func ExampleGoroutine2() {

	postpone("A goroutine starts a new thread.", 5*time.Second)
	fmt.Println("Let’s hope the news will published before I leave.")

	// Wait for the news to be published.
	time.Sleep(10 * time.Second)

	fmt.Println("Ten seconds later: I’m leaving now.")

	// Output:
}

func postpone(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
	}()
}

// which is used to broadcast a message when the text has been published.
func postpone2(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{}) // This clearly indicates that the channel will only be used for signalling, not for passing data.
	go func() {
		time.Sleep(delay)
		fmt.Println(text)
		close(ch)
	}()
	return ch
}

func ExampleGoroutine3() {
	wait := postpone2("A goroutine starts a new thread. 2", 5*time.Second)
	<-wait // Block until the text has been published.
	// Output:
}

func ExampleGoroutine4() {
	// look: ExampleWaitGroup1()
	// look: ExampleChannel3()
	for i := 0; i < 10; i++ {
		go func(a, b int) {
			c := a + b
			fmt.Printf("%d + %d = %d\n", a, b, c)
		}(i, i+1)
	}
	time.Sleep(time.Second * 2)
	// Output:
}

func ExampleGoroutine5() {
	number := func() chan int {
		ch := make(chan int)
		go func() {
			n := 1
			for {
				select {
				case ch <- n:
					n++
				case <-ch:
					return
				}
			}
		}()
		return ch
	}()
	fmt.Println(<-number)
	fmt.Println(<-number)
	close(number)
	// Output:
}
