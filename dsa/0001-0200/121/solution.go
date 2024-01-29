// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

// Example 1:

// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
// Example 2:

// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.

package main

import (
	"fmt"
	"math"
)

// Try 1
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	max := 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		min = int(math.Min(float64(min), float64(prices[i])))
		max = int(math.Max(float64(max), float64(prices[i]-min)))
	}

	return max
}

// Try 2
func maxProfit2(prices []int) int {
	buy := 0
	sell := 1
	max := 0

	for sell < len(prices) {
		if prices[sell] > prices[buy] {
			profit := prices[sell] - prices[buy]
			max = int(math.Max(float64(max), float64(profit)))
		} else {
			buy = sell
		}
		sell++
	}
	return max
}

func main() {
	result1 := maxProfit([]int{7, 1, 5, 3, 6, 4})
	result2 := maxProfit2([]int{7, 6, 4, 3, 1})

	fmt.Println(result1)
	fmt.Println(result2)
}
