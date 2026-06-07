package main

import "testing"

type testCase struct {
	word1    string
	word2    string
	expected int
}

var tests = []testCase{
	{
		word1:    "horse",
		word2:    "ros",
		expected: 3,
	}, {
		word1:    "intention",
		word2:    "execution",
		expected: 5,
	},
}

func TestMinDistance(t *testing.T) {
	for _, test := range tests {
		actual := minDistance(test.word1, test.word2)
		if actual != test.expected {
			t.Errorf("minDistance(%v, %v) expected %v, actual %v", test.word1, test.word2, test.expected, actual)
		}
	}
}
