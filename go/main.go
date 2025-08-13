package main

import "fmt"

func main() {
	roman1 := "III"
	roman2 := "LVIII"
	roman3 := "MCMXCIV"

	out1 := romanToInt(roman1)
	out2 := romanToInt(roman2)
	out3 := romanToInt(roman3)

	fmt.Println("1: ", out1)
	fmt.Println("2: ", out2)
	fmt.Println("3: ", out3)

}
