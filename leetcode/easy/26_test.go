package easy

import "fmt"

func Example26() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	// Output:
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	index, length := 0, len(nums)
	for i := 0; i < length-1; i++ {
		if nums[index] == nums[index+1] {
			nums = append(nums[:index], nums[index+1:]...)
			continue
		}
		index++
	}
	return len(nums)
}
