package binarySearch

import "fmt"

func FindPosition() {
	nums := []int{1, 3, 5, 6}
	target := 6
	left, right := 0, len(nums)-1
	index := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			index = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("index is:", index)
}
