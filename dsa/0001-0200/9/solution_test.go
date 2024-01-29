package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{1221, true},
	}

	// Run tests
	for _, test := range tests {
		result := isPalindrome(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %t, but got %t", test.input, test.expected, result)
		}
	}
}
