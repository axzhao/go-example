package unsafe

import (
	"fmt"
	"reflect"
	"unsafe"
)

func ExampleUnsafe() {
	s := "hello, world"
	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]
	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)   // 12
	fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len) // 5
	fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len) // 5
	// Output:
}

// Go 的指针是不支持指针运算和转换
// func ExampleUnsafe2() {
// 	num := 5
// 	numPointer := &num

// 	flnum := (*float32)(numPointer)
// 	fmt.Println(flnum)

// 	// Output:
// }

func ExampleUnsafe3() {
	num := 5
	numPointer := &num

	flnum := (*float32)(unsafe.Pointer(numPointer))
	fmt.Println(flnum)
	// Output:
}

// 结构体的成员变量在内存存储上是一段连续的内存
// 结构体的初始地址就是第一个成员变量的内存地址
// 基于结构体的成员地址去计算偏移量。就能够得出其他成员变量的内存地址

func ExampleUnsafe4() {

	type Num struct {
		i string
		j int64
	}

	n := Num{i: "123", j: 1}
	nPointer := unsafe.Pointer(&n)

	niPointer := (*string)(unsafe.Pointer(nPointer))
	*niPointer = "中文"

	// 这里存在一个问题，uintptr 类型是不能存储在临时变量中的。因为从 GC 的角度来看，uintptr 类型的临时变量只是一个无符号整数，并不知道它是一个指针地址
	// 因此当满足一定条件后，ptr 这个临时变量是可能被垃圾回收掉的
	// ptr := uintptr(nPointer)
	// njPointer := (*int64)(unsafe.Pointer(ptr + unsafe.Offsetof(n.j)))
	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
	*njPointer = 2

	fmt.Printf("n.i: %s, n.j: %d", n.i, n.j)

	// Output:
}
