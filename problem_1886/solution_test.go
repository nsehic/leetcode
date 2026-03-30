package main

import "testing"

type testCase struct {
	matrix, target [][]int
	expected       bool
}

var tests = []testCase{
	{
		matrix:   [][]int{{0, 1}, {1, 0}},
		target:   [][]int{{1, 0}, {0, 1}},
		expected: true,
	}, {
		matrix:   [][]int{{0, 1}, {1, 0}},
		target:   [][]int{{1, 1}, {1, 0}},
		expected: false,
	},
}

func TestFindRotation(t *testing.T) {
	for _, test := range tests {
		actual := findRotation(test.matrix, test.target)
		if actual != test.expected {
			t.Errorf("findRotation(%v, %v) expected %v, actual %v", test.matrix, test.target, test.expected, actual)
		}
	}
}
