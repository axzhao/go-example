package error

import "fmt"

func ExamplePanic() {
	n := foo2()
	fmt.Println("main received", n)

	// Output:
}

func foo2() (m int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic error", err)
			m = 2
		}
	}()
	m = 1
	panic("foo: fail")
	m = 3
	return m
}
