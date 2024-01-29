package main

import (
	"fmt"
	"math"
	"testing"
)

// Try 1
func mergeAlternatively(word1 string, word2 string) string {
	max := math.Max(float64(len(word1)), float64(len(word2)))
	results := ""

	for i := 0; i < int(max); i++ {
		if i < len(word1) {
			results += string(word1[i])
		}

		if i < len(word2) {
			results += string(word2[i])
		}
	}

	return results
}

// Try 2
func mergeAlternatively2(word1 string, word2 string) string {
	max := int(math.Max(float64(len(word1)), float64(len(word2))))
	res := make([]byte, 0, max*2)

	for i := range make([]struct{}, max) {

		if i < len(word1) {
			res = append(res, word1[i])
		}

		if i < len(word2) {
			res = append(res, word2[i])
		}

	}

	return string(res)
}

func BenchmarkMergeByteSlice(b *testing.B) {
	word1 := "hello"
	word2 := "world"

	for i := 0; i < b.N; i++ {
		_ = mergeAlternatively(word1, word2)
	}
}

func BenchmarkMergeStringConcat(b *testing.B) {
	word1 := "hello"
	word2 := "world"

	for i := 0; i < b.N; i++ {
		_ = mergeAlternatively2(word1, word2)
	}
}

func main() {

	fmt.Println(mergeAlternatively2("abc", "pqr"))

}
