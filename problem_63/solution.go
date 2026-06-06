package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1

	for row := range dp {
		for col := range dp[row] {
			if obstacleGrid[row][col] == 1 {
				dp[row][col] = 0
				continue
			}
			paths := 0

			if row > 0 {
				paths += dp[row-1][col]
			}

			if col > 0 {
				paths += dp[row][col-1]
			}

			dp[row][col] = dp[row][col] + paths
		}
	}

	return dp[m-1][n-1]
}
