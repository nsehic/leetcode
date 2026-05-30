package main

import "slices"

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		max := 0
		for k := range i {
			if nums[k] < nums[i] {
				if dp[k] > max {
					max = dp[k]
				}
			}
		}
		dp[i] = 1 + max
	}

	return slices.Max(dp)
}
