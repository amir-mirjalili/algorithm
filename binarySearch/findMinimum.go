package binarySearch

import "fmt"

func FindMinimum() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[left] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	fmt.Println("FindMinimum:", nums[left])
}
