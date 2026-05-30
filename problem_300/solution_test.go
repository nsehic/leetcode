package main

import "testing"

type testCase struct {
	input    []int
	expected int
}

var tests = []testCase{
	{
		input:    []int{10, 9, 2, 5, 3, 7, 101, 18},
		expected: 4,
	}, {
		input:    []int{0, 1, 0, 3, 2, 3},
		expected: 4,
	},
	{
		input:    []int{7, 7, 7, 7, 7, 7, 7},
		expected: 1,
	},
}

func TestLengthOfLIS(t *testing.T) {
	for _, test := range tests {
		actual := lengthOfLIS(test.input)
		if actual != test.expected {
			t.Errorf("lengthOfLIS(%v) expected %v, actual %v", test.input, test.expected, actual)
		}
	}
}
