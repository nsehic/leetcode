package main

import "testing"

type testCase struct {
	arr      []int
	start    int
	expected bool
}

var tests = []testCase{
	{
		arr:      []int{4, 2, 3, 0, 3, 1, 2},
		start:    5,
		expected: true,
	}, {
		arr:      []int{4, 2, 3, 0, 3, 1, 2},
		start:    0,
		expected: true,
	}, {
		arr:      []int{3, 0, 2, 1, 2},
		start:    2,
		expected: false,
	},
}

func TestCanReach(t *testing.T) {
	for _, test := range tests {
		actual := canReach(test.arr, test.start)
		if actual != test.expected {
			t.Errorf("canReach(%v, %v) expected %v, actual %v", test.arr, test.start, test.expected, actual)
		}
	}
}
