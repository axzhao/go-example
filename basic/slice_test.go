package basic

import "fmt"

func ExampleSlice() {
	var a []int
	fmt.Println(a == nil, len(a), cap(a), append(a, 1))
	b := []int{}
	fmt.Println(b == nil, len(b), cap(b), append(b, 1))
	c := make([]int, 0)
	fmt.Println(c == nil, len(c), cap(c), append(c, 1))
	d := make([]int, 5)
	fmt.Println(d == nil, len(d), cap(d), append(d, 1))
	e := make([]int, 0, 5)
	fmt.Println(e == nil, len(e), cap(e), append(e, 1))
	// Output:
}

func ExampleSliceInsert() {
	x := 1
	i := 5
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a = append(a[:i], append([]int{x}, a[i:]...)...) // 临时切片
	// a = append(a[:i], append([]int{1, 2, 3}, a[i:]...)...) // 在第i个位置插入切片
	fmt.Println(a)

	a = append(a, 0)     // 切片扩展1个空间
	copy(a[i+1:], a[i:]) // a[i:]向后移动1个位置
	a[i] = x             // 设置新添加的元素
	// a = append(a, x...)       // 为x切片扩展足够的空间
	// copy(a[i+len(x):], a[i:]) // a[i:]向后移动len(x)个位置
	// copy(a[i:], x)            // 复制新添加的切片
	fmt.Println(a)
	// Output:
}
