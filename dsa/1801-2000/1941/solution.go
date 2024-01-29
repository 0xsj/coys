package main

import (
	"fmt"
)

// Given a string s, return true if s is a good string, or false otherwise.

// A string s is good if all the characters that appear in s have the same number of occurrences (i.e., the same frequency).

// Example 1:

// Input: s = "abacbc"
// Output: true
// Explanation: The characters that appear in s are 'a', 'b', and 'c'. All characters occur 2 times in s.
// Example 2:

// Input: s = "aaabb"
// Output: false
// Explanation: The characters that appear in s are 'a' and 'b'.
// 'a' occurs 3 times while 'b' occurs 2 times, which is not the same number of times.

// Try 1
func areOccurencesEqual(s string) bool {
	occurences := make(map[rune]int)

	for _, char := range s {
		occurences[char]++
	}

	firstCharCount := 0
	for _, count := range occurences {
		firstCharCount = count
		break
	}

	for _, count := range occurences {
		if count != firstCharCount {
			return false
		}
	}

	return true
}

// Try 2
func areOccurencesEqual2(s string) bool {
	charCount := 0

	for i, char := range s {
		if i == 0 {
			charCount = 1
		} else if char != rune(s[i-1]) {
			return false
		} else {
			charCount++
		}
	}

	return true
}

func main() {
	result := areOccurencesEqual("hello")
	fmt.Println(result)
}
