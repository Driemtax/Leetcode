package main

import "fmt"

func main() {
	x1 := 2.00000
	n1 := 10

	x2 := 2.10000
	n2 := 3

	x3 := 2.00000
	n3 := -2

	r1 := myPow(x1, n1)
	r2 := myPow(x2, n2)
	r3 := myPow(x3, n3)

	fmt.Printf("Result 1: %f \n", r1)
	fmt.Printf("Result 2: %f \n", r2)
	fmt.Printf("Result 3: %f \n", r3)

}
