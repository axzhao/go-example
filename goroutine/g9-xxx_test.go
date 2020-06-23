package goroutine

import (
	"fmt"
	"time"
)

type batchRequest struct {
	key     int
	channel chan int
}

func Example() {

	// load
	input := make(chan *batchRequest, 5)
	go func() {
		fmt.Println("1")
		c := make(chan int, 1)
		input <- &batchRequest{1, c}
	}()
	// load
	go func() {
		fmt.Println("2")
		c := make(chan int, 2)
		input <- &batchRequest{2, c}
	}()
	// load
	go func() {
		fmt.Println("3")
		c := make(chan int, 3)
		input <- &batchRequest{3, c}
	}()

	// sleeper
	go func() {
		select {
		case <-time.After(9 * time.Second):
			fmt.Println("close")
			close(input)
		}
	}()

	// batch
	go func() {
		fmt.Println("zuse")
		for i := range input {
			fmt.Println("zhuijia", i)
		}
		fmt.Println("pass")
		fmt.Println("hebingjisuan")
	}()

	select {
	case <-time.After(10 * time.Second):
		return
	}

	// Output:
}
