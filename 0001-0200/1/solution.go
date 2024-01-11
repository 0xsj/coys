package main

import "fmt"

/*
Brute force
*/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// hashmap

func twoSum2(nums []int, target int) []int {
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

func main() {
	twoSum([]int{2, 5, 7, 9, 19}, 9)
	fmt.Println(twoSum2([]int{2, 5, 7, 9, 19}, 9))
}
