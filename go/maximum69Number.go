package main

func maximum69Number(num int) int {
	originalNum := num
	maxNumber := 0
	changed := false
	digits := []int{}

	// get all digits
	for num > 0 {
		digit := num % 10
		digits = append([]int{digit}, digits...)
		num /= 10
	}

	// swap first 6 for 9
	for i, d := range digits {
		if d == 6 {
			digits[i] = 9
			changed = true
			break
		}
	}

	if !changed {
		return originalNum
	}

	weight := 1
	// form int again
	for i := len(digits) - 1; i >= 0; i-- {
		maxNumber += digits[i] * weight
		weight *= 10
	}

	return maxNumber
}
