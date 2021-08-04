package lib

import (
	"math"
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

func CountLargedUniqueSubstring(input string) int {
	// visited array that be index by rune (int32)
	largest := 0
	for index, currentChar := range input {
		// set start char of sequence as visted
		count := 1
		// create map of seen runes
		visited := map[rune]bool{currentChar: true}

		// iterate until the end of the string starting at index
		for _, nextChar := range input[index+1:] {
			// check if next char in seqeunce has been previously visited
			if _, ok := visited[nextChar]; ok {
				break
			}
			visited[nextChar] = true
			count++
		}

		// here we set the max unique substring
		largest = int(
			math.Max(
				float64(largest), float64(count)))
	}
	return largest
}
