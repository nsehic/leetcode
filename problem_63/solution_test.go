package main

import "testing"

type testCase struct {
	grid     [][]int
	expected int
}

var tests = []testCase{
	{
		grid:     [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		expected: 2,
	}, {
		grid:     [][]int{{0, 1}, {0, 0}},
		expected: 1,
	},
}

func TestUniquePathsWithObstacles(t *testing.T) {
	for _, test := range tests {
		actual := uniquePathsWithObstacles(test.grid)
		if actual != test.expected {
			t.Errorf("uniquePathsWithObstacles(%v) expected %v, actual %v", test.grid, test.expected, actual)
		}
	}
}
