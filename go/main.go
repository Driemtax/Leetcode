package main

import "fmt"

func main() {
	input := []int{4, 8, 64}
	for _, i := range input {
		r := mySqrt(i)
		fmt.Println(r)
	}

}
