package goroutine

import (
	"fmt"
	"time"
)

// block: nil / full

// c := make(chan int, 2) // 缓存区大小为2
// c <- 1  // 立即返回
// c <- 2  // 立即返回
// c <- 3  // 阻塞直至其他goroutine调用 <- c 接收1

func f() error {
	// time.Sleep(5 * time.Second)
	// return nil
	return fmt.Errorf("%s", "2")
}

func channel1() error {
	errc := make(chan error, 2)
	var timer *time.Timer
	// 超时
	timer = time.AfterFunc(3*time.Second, func() {
		errc <- fmt.Errorf("%s", "1")
	})
	go func() {
		// 失败
		err := f()
		if timer != nil {
			timer.Stop()
		}
		errc <- err
	}()
	return <-errc
}

func ExampleChannel1() {
	// 阻塞等待结果，如果失败或超时，关闭底层连接
	if err := channel1(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("haha")
	time.Sleep(5 * time.Second)
	// Output:
}

func ExampleChannel2() {

	ch := make(chan string, 2)
	ch <- "hello"
	fmt.Println(<-ch)
	// fmt.Println(<-ch) block

	ch2 := make(chan string, 2)
	go func() {
		ch2 <- "Hello!"
		close(ch2) // The close function records that no more values will be sent on a channel.
	}()

	fmt.Println(<-ch2) // Print "Hello!".
	fmt.Println(<-ch2) // Print the zero value "" without blocking.

	// Output:
}

func cal(a int, b int, Exitchan chan bool) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	time.Sleep(time.Second * 2)
	Exitchan <- true
}

func ExampleChannel3() {
	Exitchan := make(chan bool, 10) //声明并分配管道内存
	for i := 0; i < 10; i++ {
		go cal(i, i+1, Exitchan)
	}
	// 使用range循环管道，如果管道未关闭会引发deadlock错误。
	// 如果采用for死循环已经关闭的管道，当管道没有数据时候，读取的数据会是管道的默认值，并且循环不会退出。
	for j := 0; j < 10; j++ {
		<-Exitchan //取信号数据，如果取不到则会阻塞
	}
	close(Exitchan) // 关闭管道
	// Output:
}
