package main

import "fmt"

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for i != nums[i] {
			//fmt.Println(i, nums)
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			temp := 0
			temp = nums[i]
			nums[i], nums[temp] = nums[temp], nums[i]
		}
	}
	return -1
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
		result = findRepeatNumber(numsList)
		fmt.Println(result)
	}
}
