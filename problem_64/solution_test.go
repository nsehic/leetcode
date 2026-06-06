package main

import "testing"

type testCase struct {
	grid     [][]int
	expected int
}

var tests = []testCase{
	{
		grid:     [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
		expected: 7,
	}, {
		grid:     [][]int{{1, 2, 3}, {4, 5, 6}},
		expected: 12,
	},
}

func TestMinimumPathSum(t *testing.T) {
	for _, test := range tests {
		actual := minPathSum(test.grid)
		if actual != test.expected {
			t.Errorf("minPathSum(%v) expected %v, actual %v", test.grid, test.expected, actual)
		}
	}
}
