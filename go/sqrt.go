package main

func mySqrt(x int) int {
	result := 0

	for (result+1024)*(result+1024) < x {
		result += 1024
	}

	for {
		sqrt := result * result

		if sqrt == x {
			return result
		}

		if sqrt > x {
			return result - 1
		}

		result += 1
	}
}
