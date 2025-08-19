package main

// Still kind of proud of my first implementation thats why i left it in. Sadly its of complexy O(n^2) in worst case and leetcode checks for that
// func zeroFilledSubarray(nums []int) int64 {
// 	zeros := []int{}

// 	// scan for zeros
// 	for i, v := range nums {
// 		if v == 0 {
// 			zeros = append(zeros, i)
// 		}
// 	}

// 	if len(zeros) == 0 {
// 		return 0
// 	}

// 	// count subarrays
// 	result := len(zeros)

// 	for i := 0; i < len(zeros); i++ {
// 		counter := 1
// 		for i+counter < len(zeros) && zeros[i+counter]-zeros[i] == counter {
// 			counter++
// 			result++
// 		}
// 	}

// 	return int64(result)
// }

// This is much more efficient with O(n)
func zeroFilledSubarray(nums []int) int64 {
	var result int64 = 0
	var count int64 = 0

	for _, v := range nums {
		if v == 0 {
			count++
			result += count
		} else {
			count = 0
		}
	}
	return result
}
