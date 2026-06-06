package main

import "testing"

type testCase struct {
	input    []int
	expected int
}

var tests = []testCase{
	{
		input:    []int{2, 3, 1, 1, 4},
		expected: 2,
	}, {
		input:    []int{2, 3, 0, 1, 4},
		expected: 2,
	},
}

func TestJumpDP(t *testing.T) {
	for _, test := range tests {
		actual := jumpDP(test.input)
		if actual != test.expected {
			t.Errorf("jumpDP(%v) expected %v, actual %v", test.input, test.expected, actual)
		}
	}
}

func TestJumpGreedy(t *testing.T) {
	for _, test := range tests {
		actual := jumpGreedy(test.input)
		if actual != test.expected {
			t.Errorf("jumpGreedy(%v) expected %v, actual %v", test.input, test.expected, actual)
		}
	}
}
