package main

func largestGoodInteger(num string) string {
	result := ""

	for i := 2; i < len(num); i++ {
		if num[i-1] == num[i] && num[i-2] == num[i] {
			if len(result) == 0 || num[i] > result[0] {
				result = num[i-2 : i+1]
			}
		}
	}

	return result
}
