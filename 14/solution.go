package main

import "fmt"

// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

// Constraints:

// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters.

// Try 1
func longestCommonPrefix(strs []string) string {
	// base case
	if len(strs) == 0 {
		return ""
	}

	first := strs[0]

	for i := 0; i < len(first); i++ {
		current := first[i]

		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != current {
				return first[:i]
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"fldog", "flow", "flight"}))
}
