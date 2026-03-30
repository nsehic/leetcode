package main

import "testing"

type testCase struct {
	s1, s2   string
	expected bool
}

var tests = []testCase{
	{
		s1:       "abcd",
		s2:       "cdab",
		expected: true,
	}, {
		s1:       "abcd",
		s2:       "dacb",
		expected: false,
	},
}

func TestCanBeEqual(t *testing.T) {
	for _, test := range tests {
		actual := canBeEqual(test.s1, test.s2)
		if actual != test.expected {
			t.Errorf("canBeEqual(%v, %v) expected %v, actual %v", test.s1, test.s2, test.expected, actual)
		}
	}
}
