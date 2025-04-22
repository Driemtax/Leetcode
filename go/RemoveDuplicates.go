package main

import (
	"slices"
)

func removeDuplicates(nums []int) int {
	uniqueIndex := 0
	uniquenNums := []int{}
	for i := 0; i < len(nums); i++ {
		if !(slices.Contains(uniquenNums, nums[i])) {
			nums[uniqueIndex] = nums[i]
			uniquenNums = append(uniquenNums, nums[i])
			uniqueIndex++
		}
	}

	return uniqueIndex
}
