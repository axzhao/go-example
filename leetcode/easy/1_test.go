package easy

import "fmt"

func ExampleTwoSum() {

	nums := []int{-2, 7, 11, 15}
	target := 5

	fmt.Println(twoSum(nums, target))

	// Output:

}

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		panic("")
	}
	tmp := make(map[int]int, len(nums))
	for i, v := range nums {
		if vv, ok := tmp[v]; ok {
			return []int{vv, i}
		}
		tmp[target-v] = i
	}
	return []int{}
}
