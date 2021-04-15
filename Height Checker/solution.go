package main

import (
	"fmt"
)

func sortArray(res []int) {
	for j := 0; j < len(res); j++ {
		for k := 1; k < len(res)-j; k++ {
			if res[k] < res[k-1] {
				a := res[k]
				res[k] = res[k-1]
				res[k-1] = a
			}
		}
	}
}

func heightChecker(heights []int) int {
	weights := make([]int, len(heights))
	copy(weights, heights)
	sortArray(heights)
	count := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != weights[i] {
			count = count + 1
		}
	}
	return count
}

func main() {
	for {
		var numsSize, inputNum int
		fmt.Scan(&numsSize)
		var numsList []int
		for j := 0; j < numsSize; j++ {
			fmt.Scan(&inputNum)
			numsList = append(numsList, inputNum)
		}
		result := 0
		result = heightChecker(numsList)
		fmt.Println(result)
	}
}
