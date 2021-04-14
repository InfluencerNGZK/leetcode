package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minValue := math.MaxInt64
	maxValue := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minValue {
			minValue = prices[i]
		} else if prices[i]-minValue > maxValue {
			maxValue = prices[i] - minValue
		}
	}
	return maxValue
}

func main() {
	//examplePrice := []int{7, 1, 5, 3, 6, 4}
	var input, nums int
	fmt.Scan(&nums)
	fmt.Println(nums)
	var examplePrice []int
	for j := 0; j < nums; j++ {
		fmt.Scan(&input)
		examplePrice = append(examplePrice, input)
	}
	result := 0
	result = maxProfit(examplePrice)
	fmt.Println(result)
}
