package main

func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)

	if n == 1 {
		return mat[0][0] == target[0][0]
	}

	matches90 := true
	matches180 := true
	matches270 := true
	matches360 := true

	for row := range mat {
		for column := range mat[row] {
			if matches90 {
				// calculation for 90deg rotation
				// target row = mat column
				// target column = n - 1 - mat row
				matches90 = mat[row][column] == target[column][n-1-row]
			}

			if matches180 {
				// calculation for 180deg rotation
				// target row = n - 1 - mat row
				// target column = n - 1 - mat col
				matches180 = mat[row][column] == target[n-1-row][n-1-column]
			}

			if matches270 {
				// calculation for 270deg rotation
				// target row = n - 1 - mat column
				// target column = mat row
				matches270 = mat[row][column] == target[n-1-column][row]
			}

			if matches360 {
				// calculation for 360deg rotation
				matches360 = mat[row][column] == target[row][column]
			}

			// If none of the rotations are matches at the end, return false
			if !matches90 && !matches180 && !matches270 && !matches360 {
				return false
			}
		}
	}

	// Otherwise, at least one rotation matches, so return true
	return true
}
