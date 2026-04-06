package main

import "testing"

type testCase struct {
	input    string
	expected bool
}

var tests = []testCase{
	{
		input:    "UD",
		expected: true,
	}, {
		input:    "LL",
		expected: false,
	},
}

func TestJudgeCircle(t *testing.T) {
	for _, test := range tests {
		actual := judgeCircle(test.input)
		if actual != test.expected {
			t.Errorf("judgeCircle(%v) expected %v, actual %v", test.input, test.expected, actual)
		}
	}
}
