package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 5, 6, 7, 8}
	target := 4
	index := searchInsert(nums, target)
	fmt.Println(index)
}
