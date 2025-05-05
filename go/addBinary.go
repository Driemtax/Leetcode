package main

func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	} else if len(b) == 0 {
		return a
	}

	if len(a) > len(b) {
		a, b = b, a
	}

	diff := len(b) - len(a)
	for i := 0; i < diff; i++ {
		a = "0" + a
	}

	carry := 0
	result := ""

	for i := len(a) - 1; i >= 0; i-- {
		sum := carry

		if a[i] == '1' {
			sum++
		}
		if b[i] == '1' {
			sum++
		}

		switch sum {
		case 0: // a=0 && b=0
			result = "0" + result
			carry = 0
		case 1: // a=1 || b=1 || carry=1
			result = "1" + result
			carry = 0
		case 2: // a=1 && b=1 && carry=0
			result = "0" + result
			carry = 1
		case 3: // a=1 && b=1 && carry=1
			result = "1" + result
			carry = 1
		}
	}

	if carry == 1 {
		result = "1" + result
	}

	return result
}
