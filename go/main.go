package main

import "fmt"

func main() {
	// Beispiel 1
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Original Matrix 1:")
	printMatrix(matrix1)
	rotate(matrix1)
	fmt.Println("Rotierte Matrix 1:")
	printMatrix(matrix1) // Erwartet: [[7,4,1],[8,5,2],[9,6,3]]

	fmt.Println("---")

	// Beispiel 2
	matrix2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	fmt.Println("Original Matrix 2:")
	printMatrix(matrix2)
	rotate(matrix2)
	fmt.Println("Rotierte Matrix 2:")
	printMatrix(matrix2) // Erwartet: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
}
