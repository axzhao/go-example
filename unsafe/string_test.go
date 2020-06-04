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
