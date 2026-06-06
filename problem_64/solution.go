package main

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for row := range dp {
		for col := range dp[row] {
			var minimumSum int

			if row > 0 && col > 0 {
				minimumSum = min(dp[row-1][col], dp[row][col-1])
			} else if row > 0 {
				minimumSum = dp[row-1][col]
			} else if col > 0 {
				minimumSum = dp[row][col-1]
			}

			dp[row][col] = minimumSum + grid[row][col]
		}
	}

	return dp[m-1][n-1]
}
