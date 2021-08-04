package lib

import (
	"testing"
)

type example struct {
	Input    string
	Expected int
}

func TestSubstringLength(t *testing.T) {
	examples := []example{
		{"", 0},
		{"bbbbb", 1},
		{"abc", 3},
		{"abcabcbb", 3},
		{"pwwkew", 3},
		{"abcbdefgf", 6},
		{"abcabcdefghijklmnopqrstuvwxyzxyz", 26},
	}

	for _, e := range examples {
		actual := CountLargedUniqueSubstring(e.Input)

		if actual != e.Expected {
			t.Errorf("expected output %d, actual %d", e.Expected, actual)
		}
	}
}
