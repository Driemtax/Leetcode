package main

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		n := len(matrix[i]) - 1
		// found the right row
		if matrix[i][n] >= target {
			// now search row from back to front
			for j := n; j >= 0; j-- {
				element := matrix[i][j]
				// 3 cases for element (matrx[i][j]):
				// 1. element == target -> target found
				// 2. element < target -> target not in matrix since last element was bigger than target
				// 3. element > target -> continue to search
				if element == target {
					return true
				} else if element < target {
					return false
				} else {
					continue
				}
			}

			// if we reach this, the target was not in the correct row therefore not in the matrix
			return false

		} else {
			continue // look in the next row
		}
	}

	// means we didnt find target
	return false
}
