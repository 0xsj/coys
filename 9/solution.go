package main

import (
	"fmt"
	"math"
	"strconv"
)

// Try 1
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	reverse := 0

	for x > reverse {
		reverse = reverse*10 + (x % 10)
		x = int(math.Floor(float64(x) / 10))
	}

	return x == reverse || x == int(math.Floor(float64(reverse)/10)) // Corrected type conversion
}

// Try 2
func isPalindrome2(x int) bool {
	numStr := strconv.FormatInt(int64(x), 10)

	runes := []rune(numStr)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	reversedString := string(runes)

	return numStr == reversedString
}

// Try 3 simplified
func isPalindrome3(x int) bool {
	if x < 0 {
		return false
	}

	original, reversed := x, 0

	for x > 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}

	return original == reversed
}

func main() {
	fmt.Println(isPalindrome2(11211))
	fmt.Println(isPalindrome2(121))
	fmt.Println(isPalindrome2(-11211))
}
