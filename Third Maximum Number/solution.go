package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for k := 0; k < len(nums); k++ {
		if nums[k] >= a {
			a = nums[k]
		}
	}
	for j := 0; j < len(nums); j++ {
		if nums[j] >= b && nums[j] != a {
			b = nums[j]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] >= c && nums[i] != b && nums[i] != a {
			c = nums[i]
		}
	}
	if c != math.MinInt64 {
		return c
	} else {
		return a
	}
}

func main() {
	for {
		var numsSize, inputNum, result int
		fmt.Scan(&numsSize)
		var numsExample []int
		for k := 0; k < numsSize; k++ {
			fmt.Scan(&inputNum)
			numsExample = append(numsExample, inputNum)
		}
		result = thirdMax(numsExample)
		fmt.Println(result)
	}
}
