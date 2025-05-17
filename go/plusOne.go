package main

func plusOne(digits []int) []int {
	carry := 0
	if digits[len(digits)-1] == 9 {
		digits[len(digits)-1] = 0
		carry = 1
	} else {
		digits[len(digits)-1] += 1
	}

	if carry == 0 {
		return digits
	}

	for i := len(digits) - 2; i >= 0; i-- {
		if carry == 0 {
			return digits
		}
		sum := digits[i] + carry
		carry = sum / 10
		digits[i] = sum % 10
	}

	if carry == 0 {
		return digits
	}
	newDigits := []int{}
	if carry != 0 {
		newDigits = append([]int{carry}, digits...)
	}

	return newDigits
}
