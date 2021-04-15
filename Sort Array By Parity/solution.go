package main

import "fmt"

func sortArrayByParity(A []int) []int {
	evenCount := 0
	oddCount := 0
	B := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			B[evenCount] = A[i]
			evenCount++
		} else {
			B[len(A)-1-oddCount] = A[i]
			oddCount++
		}
	}
	return B
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
		resultList = sortArrayByParity(numsList)
		fmt.Println(resultList)
	}
}
