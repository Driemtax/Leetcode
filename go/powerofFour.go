package main

func isPowerOfFour(n int) bool {
	// 1. n must be positive since base is positive
	// 2. every power of 4 is also a power of 2 and therefore only has one bit set to 1
	// 3. not every power of is a power of 4, so check that only 1,3,5,... bit is set to 1
	return n > 0 && (n&(n-1)) == 0 && (n&0xAAAAAAAA == 0)
}
