package main

import "fmt"

func main() {
	matrix1 := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target1 := 3

	result1 := searchMatrix(matrix1, target1)
	fmt.Println("Result1: ", result1)

	target2 := 13

	result2 := searchMatrix(matrix1, target2)
	fmt.Println("Result2: ", result2)
}
