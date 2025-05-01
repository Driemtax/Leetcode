package main

import "fmt"

func searchInsert(nums []int, target int) int {
	index := binarySearch(nums, target)

	return index

}

func binarySearch(nums []int, target int) int { // nums = [1,2,3,5,6,7,8], target = 4, len=7
	low, high := 0, len(nums)-1

	for low < high {
		fmt.Printf("Low: %d, High: %d \n", low, high)
		mid := (low + high) / 2
		fmt.Printf("Mid: %d \n", mid)

		// check if target is mid
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target { // target is bigger then mid
			low = mid + 1
		} else { // target is smaller then mid
			high = mid - 1
		}
	}

	fmt.Printf("Low: %d, High: %d \n", low, high)

	// Element not found
	return low
}
