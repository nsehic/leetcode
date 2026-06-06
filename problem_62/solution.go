package main

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1

	for row := range dp {
		for col := range dp[row] {
			fromTop := 0
			fromLeft := 0

			if row > 0 {
				fromTop = dp[row-1][col]
			}

			if col > 0 {
				fromLeft = dp[row][col-1]
			}

			dp[row][col] = dp[row][col] + (fromTop + fromLeft)
		}
	}

	return dp[m-1][n-1]
}
