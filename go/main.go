package main

import "fmt"

func main() {
	roman1 := "6777133339"
	roman2 := "2300019"
	roman3 := "42352338"

	out1 := largestGoodInteger(roman1)
	out2 := largestGoodInteger(roman2)
	out3 := largestGoodInteger(roman3)

	fmt.Println("1: ", out1)
	fmt.Println("2: ", out2)
	fmt.Println("3: ", out3)

}
