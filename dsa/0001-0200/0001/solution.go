package main

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

import (
	"fmt"
	"time"
)

type TwoSum struct {
	nums   []int
	target int
}

/*
Brute force
*/

func (s *TwoSum) brute_force() []int {
	nums, target := s.nums, s.target

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func (s *TwoSum) hashmap() []int {
	nums, target := s.nums, s.target
	hashmap := make(map[int]int, 0)
	for i, num := range nums {
		complement := target - num
		if index, found := hashmap[complement]; found {
			return []int{index, i}
		}
		hashmap[num] = i
	}
	return []int{}
}

func measureTime(f func() []int, label string) {
	start := time.Now()
	result := f()
	elpased := time.Since(start)

	fmt.Printf("Algorithm: %s, Result: %v, Time: %s\n", label, result, elpased)
}

func (s *TwoSum) main() {
	measureTime(s.brute_force, "bruteforce")
	measureTime(s.hashmap, "hashmap")
	// result1 := s.brute_force()
	// result2 := s.hashmap()

}

// func main() {
// 	nums := []int{2, 7, 11, 15}
// 	target := 9

// 	solution := TwoSum{nums, target}
// 	solution.run()
// }
