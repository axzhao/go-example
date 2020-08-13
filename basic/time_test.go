package basic

import (
	"fmt"
	"time"
)

func ExampleTime() {

	f := "20060102T150405Z"
	fmt.Println(time.Now().UTC().Format(f))

	now := time.Now()
	fmt.Println(now.Add(time.Duration(700) * time.Second * -1))
	fmt.Println(now.Add(time.Duration(700) * time.Second * -1).Unix())
	fmt.Println(time.Now().Unix())

	// Output:
}

