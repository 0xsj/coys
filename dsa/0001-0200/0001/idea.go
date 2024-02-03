package main

import "fmt"

// Map applies a function to each element of a slice and returns a new slice.
func Map[A, B any](arr []A, f func(A) B) []B {
	result := make([]B, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

// Pair is a simple tuple structure to hold two values.
type Pair struct {
	First, Second int
}

// TwoSum_fp finds two numbers in an array that add up to the target value using a functional style.
func TwoSum_fp(nums []int, target int) []int {
	// Map the array to a slice of pairs (element, index)
	numsWithIndices := Map(nums, func(e int) Pair {
		return Pair{e, 0}
	})

	// Find the pair of indices whose values sum up to the target
	for i, num := range nums {
		complement := target - num

		// Filter the pairs to find a pair with the complement value
		result := FilterPairs(numsWithIndices, func(pair Pair) bool {
			return pair.First == complement
		})

		// If a valid pair is found, return the indices
		if len(result) > 0 {
			return []int{result[0].Second, i}
		}

		// Update the Second value in numsWithIndices based on the current index during Map
		numsWithIndices[i].Second = i
	}

	// If no solution is found, return an empty slice
	return []int{}
}

// FilterPairs filters the elements of a slice of pairs based on a given predicate function.
func FilterPairs(pairs []Pair, predicate func(Pair) bool) []Pair {
	var result []Pair
	for i, pair := range pairs {
		if predicate(pair) {
			pairs[i].Second = i // Update the index correctly
			result = append(result, pair)
		}
	}
	return result
}

func main() {
	// Example 1
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	result1 := TwoSum_fp(nums1, target1)
	fmt.Println(result1) // Output: [0 1]

	// Example 2
	// nums2 := []int{3, 2, 4}
	// target2 := 6
	// result2 := TwoSum(nums2, target2)
	// fmt.Println(result2) // Output: [1 2]

	// // Example 3
	// nums3 := []int{3, 3}
	// target3 := 6
	// result3 := TwoSum(nums3, target3)
	// fmt.Println(result3) // Output: [0 1]
}
