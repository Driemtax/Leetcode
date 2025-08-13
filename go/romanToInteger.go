package main

func romanToInt(s string) int {
	// IV = 4
	// IX = 9
	// XL = 40
	// XC = 90
	// CD = 400
	// CM = 900

	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		current := roman[s[i]]
		if current < prevValue {
			result -= current
		} else {
			result += current
		}

		prevValue = current
	}

	return result
}
