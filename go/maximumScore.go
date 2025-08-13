package main

func maxScore(s string) int {
	totalOnes := 0
	for _, c := range s {
		if c == '1' {
			totalOnes++
		}
	}

	maxScore := 0
	onesLeft := 0
	zerosLeft := 0

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			zerosLeft++
		} else {
			onesLeft++
		}

		score := zerosLeft + (totalOnes - onesLeft)
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}
