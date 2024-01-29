package main

import "fmt"

func romanToInt_1(s string) int {

	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && roman[s[i]] < roman[s[i+1]] {
			result -= roman[s[i]]
		} else {
			result += roman[s[i]]
		}
	}

	return result
}

func main() {
	result := romanToInt_1("III")
	fmt.Println(result)
}
