package main

func myPow(x float64, n int) float64 {
	// Edge case n == 0
	if n == 0 {
		return 1.0
	}
	// Edge case x == 1
	if x == 1.0 {
		return 1.0
	}
	// Edge case x == -1
	if x == -1.0 {
		if n%2 == 0 {
			return 1.0
		} else {
			return -1.0
		}
	}
	// Edge case x == 0
	if x == 0 {
		// The constraint is that n > 0 || x > 0
		return 0.0
	}

	var exponent int64 = int64(n) // convert to int64 in case of minInt32
	base := x

	//negative n
	if exponent < 0 {
		base = 1 / x
		exponent = -exponent
	}

	result := 1.0

	// positive n
	for exponent > 0 {
		if exponent%2 == 1 {
			result = result * base
		}

		base *= base
		exponent /= 2
	}

	return result
}
