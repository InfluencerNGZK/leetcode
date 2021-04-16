package main

import "fmt"

func findNumbers(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		countEven := 0
		example := nums[i]
		for example != 0 {
			example = example / 10
			countEven++
		}
		if countEven%2 == 0 {
			count++
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
		result = findNumbers(numsList)
		fmt.Println(result)
	}
}
