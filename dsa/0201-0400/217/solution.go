package main

import "fmt"

// try 1
func containsDuplicate(nums []int) bool {

	seen := make(map[int]int, 0)

	for _, num := range nums {
		if seen[num] > 0 {
			return true
		}

		seen[num]++
	}

	return false
}

// try 2
func containsDuplicate2(nums []int) bool {
	for i, num1 := range nums {
		for j, num2 := range nums {
			if i != j && num1 == num2 {
				return true
			}
		}
	}
	return false
}

// Try 3
func containsDuplicate3(nums []int) bool {
	set := make(map[int]struct{})

	for _, num := range nums {
		if _, seen := set[num]; seen {
			return true
		}

		set[num] = struct{}{}
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 1, 2, 3}))
	fmt.Println(containsDuplicate2([]int{1, 2, 3}))
	fmt.Println(containsDuplicate3([]int{1, 2, 3}))
}
