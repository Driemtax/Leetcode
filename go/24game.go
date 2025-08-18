package main

import "math"

func judgePoint24(cards []int) bool {
	nums := make([]float64, len(cards))
	for i, v := range cards {
		nums[i] = float64(v)
	}

	var dfs func([]float64) bool
	dfs = func(nums []float64) bool {
		if len(nums) == 1 {
			return math.Abs(nums[0]-24) < 1e-6
		}
		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums); j++ {
				if i == j {
					continue
				}
				next := []float64{}
				for k := 0; k < len(nums); k++ {
					if k != i && k != j {
						next = append(next, nums[k])
					}
				}
				a, b := nums[i], nums[j]
				for _, v := range []float64{a + b, a - b, b - a, a * b} {
					nextNums := append(next, v)
					if dfs(nextNums) {
						return true
					}
				}
				// Division muss prüfen, ob kein Zero-Division entsteht
				if b != 0 {
					nextNums := append(next, a/b)
					if dfs(nextNums) {
						return true
					}
				}
				if a != 0 {
					nextNums := append(next, b/a)
					if dfs(nextNums) {
						return true
					}
				}
			}
		}
		return false
	}
	return dfs(nums)
}
