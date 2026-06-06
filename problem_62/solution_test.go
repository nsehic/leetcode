package main

import "testing"

type testCase struct {
	m        int
	n        int
	expected int
}

var tests = []testCase{
	{
		m:        3,
		n:        7,
		expected: 28,
	}, {
		m:        3,
		n:        2,
		expected: 3,
	},
}

func TestUniquePaths(t *testing.T) {
	for _, test := range tests {
		actual := uniquePaths(test.m, test.n)
		if actual != test.expected {
			t.Errorf("uniquePaths(%v, %v) expected %v, actual %v", test.m, test.n, test.expected, actual)
		}
	}
}
