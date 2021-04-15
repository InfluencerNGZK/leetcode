package main

import "fmt"

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	numsNew := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if changeInt(nums[left]) >= changeInt(nums[right]) {
			numsNew[len(nums)-1-i] = nums[left] * nums[left]
			left++
		} else {
			numsNew[len(nums)-1-i] = nums[right] * nums[right]
			right--
		}
	}
	return numsNew
}

func changeInt(a int) int {
	if a < 0 {
		return a * -1
	} else {
		return a
	}
}

func main() {
	for {
		var numsSize, inputNum int
		fmt.Scan(&numsSize)
		var numsList, resultList []int
		for j := 0; j < numsSize; j++ {
			fmt.Scan(&inputNum)
			numsList = append(numsList, inputNum)
		}
		resultList = sortedSquares(numsList)
		fmt.Println(resultList)
	}
}
