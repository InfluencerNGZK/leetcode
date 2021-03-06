package main

import "fmt"

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	left := 0
	right := len(arr) - 1
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			break
		} else if arr[i] > arr[i-1] {
			left = left + 1
		}
	}
	for j := len(arr) - 1; j > 0; j-- {
		if arr[j] >= arr[j-1] {
			break
		} else if arr[j] < arr[j-1] {
			right = right - 1
		}
	}
	if right == left && right != 0 && left != len(arr)-1 {
		return true
	} else {
		return false
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
		var result bool
		result = validMountainArray(numsExample)
		fmt.Println(result)
	}
}
