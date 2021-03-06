package main

import "fmt"

func duplicateZeros(arr []int) {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			count = count + 1
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			count = count - 1
			if i+count+1 < len(arr) {
				arr[i+count+1] = arr[i]
			}
		}
		if i+count < len(arr) {
			arr[i+count] = arr[i]
		}
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
		duplicateZeros(numsList)
		fmt.Println(numsList)
	}
}
