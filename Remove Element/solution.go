package main

import "fmt"

func removeElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
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
		var removeNum, result int
		fmt.Scan(&removeNum)
		result = removeElement(numsList, removeNum)
		fmt.Println(result)
	}
}
