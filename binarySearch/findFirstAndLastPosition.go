package binarySearch

import "fmt"

func FindFirsAndLast() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	first := findFirst(nums, target)
	last := findLast(nums, target)
	fmt.Println("find first last:", first, last)
}

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	index := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			index = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return index
}

func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	index := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			index = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return index
}
