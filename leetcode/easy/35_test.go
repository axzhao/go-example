package easy

import "fmt"

func Example35() {
	fmt.Println(searchInsert([]int{1}, 1))

	// Output:
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	low, mid, high := 0, (len(nums)-1)/2, len(nums)-1
	if target > nums[high] {
		return len(nums)
	}
	if target < nums[low] {
		return 0
	}
	for high >= low {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			low = mid + 1
		}
		if nums[mid] > target {
			high = mid - 1
		}
		mid = (high + low) / 2
	}
	return mid + 1
}
