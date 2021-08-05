package lib

import (
	"fmt"
	"strings"
)

// Given a string, write a function that returns the length of the longest substring that does not contain repeating characters example:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Input: s = ""
// Output: 0

func ExtractLargestUniqueSubstring(input string) (largest string) {
	chars := strings.Split(input, "")
	for index, substring := range chars {
		// iterate until the end of the string starting at index
		for _, nextChar := range chars[index+1:] {
			// check if next char in seqeunce has been previously seen
			if strings.Contains(substring, nextChar) {
				break
			}
			substring = fmt.Sprintf("%s%s", substring, nextChar)
		}

		// here we set the max unique substring
		if len(substring) > len(largest) {
			largest = substring
		}
	}
	return largest
}
