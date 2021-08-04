package lib

import "fmt"

// Given a string, write a function that returns the length of the longest substring that does not contain repeating characters example:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Input: s = ""
// Output: 0

func CountLargestUniqueSubstring(input string) string {
	var largest string
	for index, currentChar := range input {
		// create map of seen runes
		substring := string(currentChar)
		seen := map[rune]bool{currentChar: true}

		// iterate until the end of the string starting at index
		for _, nextChar := range input[index+1:] {
			// check if next char in seqeunce has been previously seen
			if _, ok := seen[nextChar]; ok {
				break
			}
			substring = fmt.Sprintf("%s%s", substring, string(nextChar))
			seen[nextChar] = true
		}

		// here we set the max unique substring
		if len(substring) > len(largest) {
			largest = substring
		}
	}
	return largest
}
