package main

import "fmt"

func replaceElements(arr []int) []int {
	arrNew := make([]int, len(arr))
	for i := 0; i < len(arr)-1; i++ {
		a := 0
		a = arr[i+1]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > a {
				a = arr[j]
			}
		}
		arrNew[i] = a
	}
	arrNew[len(arr)-1] = -1
	return arrNew
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
		resultList = replaceElements(numsList)
		fmt.Println(resultList)
	}
}
