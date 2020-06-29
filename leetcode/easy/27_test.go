package easy

import "fmt"

func Example27() {
	fmt.Println(removeElement([]int{2}, 3))
	// Output:
}

func removeElement(nums []int, val int) int {
	index, length := 0, len(nums)
	for i := 0; i < length; i++ {
		if nums[index] == val {
			nums = append(nums[:index], nums[index+1:]...)
			continue
		}
		index++
	}
	return index
}
