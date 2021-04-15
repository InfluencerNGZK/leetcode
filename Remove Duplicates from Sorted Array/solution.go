package main

import "fmt"

func removeDuplicates(nums []int) int {
	count := 0
	if len(nums) != 0 {
		per := nums[0]
		for i := 0; i < len(nums); i++ {
			if per != nums[i] {
				count = count + 1
				nums[count] = nums[i]
				per = nums[i]
			}
		}
		return count + 1
	} else {
		return 0
	}
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
		result = removeDuplicates(numsList)
		fmt.Println(result, numsList[:result])
	}
}
