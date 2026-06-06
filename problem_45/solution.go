package main

import "math"

func jumpDP(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 0

	for i := 1; i < n; i++ {
		minJumps := math.MaxInt

		for j := 0; j < i; j++ {
			jump := nums[j]
			if jump >= i-j && dp[j] < minJumps {
				minJumps = dp[j]
			}
		}
		dp[i] = minJumps + 1
	}

	return dp[n-1]
}

func jumpGreedy(nums []int) int {
	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}
	}

	return jumps
}
