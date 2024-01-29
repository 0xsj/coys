// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

// Example 1:

// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
// Example 2:

// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.
// Example 3:

// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.

package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// // not working
// func is_palindrome(s string) bool {
// 	pattern := "[^a-zA-Z0-9]"
// 	re := regexp.MustCompile(pattern)

// 	clean := re.ReplaceAllString(s, "")
// 	runes := []rune(clean)
// 	left := 0
// 	right := len(runes) - 1

// 	for left < right {

// 		if runes[left] != runes[right] {
// 			return false
// 		}

// 		left += 1
// 		right -= 1
// 	}

// 	return true
// }

// try 2
func is_palindrome2(s string) bool {
	re := regexp.MustCompile("[^a-zA-Z0-9]")
	runes := []rune(re.ReplaceAllString(s, ""))
	left, right := 0, len(runes)-1

	for left < right {
		if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
			return false
		}
		left += 1
		right -= 1
	}

	return true
}

// try 3
func is_palindrome3(s string) bool {

	re := regexp.MustCompile("[^a-zA-Z0-9]")
	clean := re.ReplaceAllString(s, "")
	left, right := 0, len(clean)-1

	for left < right {

		if !strings.EqualFold(string(clean[right]), string(clean[left])) {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	fmt.Println(is_palindrome3("A man, a plan, a canal: Panama"))
	fmt.Println(is_palindrome3("race a car"))
}
