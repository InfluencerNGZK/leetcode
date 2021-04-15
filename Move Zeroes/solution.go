package main

import "fmt"

func moveZeroes(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count = count + 1
		}
	}
	for j := count; j < len(nums); j++ {
		nums[j] = 0
	}
}

func main() {
	for {
		var numsSize, inputNum int
		fmt.Scan(&numsSize)
		var numsExample []int
		for k := 0; k < numsSize; k++ {
			fmt.Scan(&inputNum)
			numsExample = append(numsExample, inputNum)
		}
		moveZeroes(numsExample)
		fmt.Println(numsExample)
	}
}
