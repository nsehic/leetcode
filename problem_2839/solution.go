package main

import "slices"

func canBeEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	runes := []rune(s1)
	allIndices := [][]int{{0, 2}, {1, 3}}

	for _, indices := range allIndices {
		i, j := indices[0], indices[1]
		runes[i], runes[j] = runes[j], runes[i]

		if string(runes) == s2 {
			return true
		}
	}

	runes = []rune(s1)
	for _, indices := range slices.Backward(allIndices) {
		i, j := indices[0], indices[1]
		runes[i], runes[j] = runes[j], runes[i]

		if string(runes) == s2 {
			return true
		}
	}

	return false
}
