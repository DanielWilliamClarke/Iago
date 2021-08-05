package lib

import (
	"testing"
)

type example struct {
	Input    string
	Expected string
}

func TestSubstringLength(t *testing.T) {
	examples := []example{
		{"", ""},
		{"bbbbb", "b"},
		{"abc", "abc"},
		{"abcabcbb", "abc"},
		{"pwwkew", "wke"},
		{"abcbdefgf", "cbdefg"},
		{"abcabcdefghijklmnopqrstuvwxyzxyz", "abcdefghijklmnopqrstuvwxyz"},
	}

	for _, e := range examples {
		actual := ExtractLargestUniqueSubstring(e.Input)
		if actual != e.Expected {
			t.Errorf("expected output %s, actual %s", e.Expected, actual)
		}
	}
}
