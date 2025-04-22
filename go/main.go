package main

import "fmt"

func main() {
	// Your code starts here
	fmt.Println("Hello, World!")
	nums := []int{3, 2, 2, 3}
	val := 3
	length := removeElement(nums, val)
	fmt.Println(length)
	fmt.Println(nums[:length])
}
