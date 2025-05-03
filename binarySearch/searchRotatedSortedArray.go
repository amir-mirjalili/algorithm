package binarySearch

import "fmt"

func SearchRotatedSortedArray() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			fmt.Println("SearchRotatedSortedArray :", mid)
			return
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
}
