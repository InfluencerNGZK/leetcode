package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}
	for k := range nums1 {
		for j := 1; j < len(nums1)-k; j++ {
			if nums1[j] < nums1[j-1] {
				a := nums1[j]
				nums1[j] = nums1[j-1]
				nums1[j-1] = a
			}
		}
	}
}

func main() {
	for {
		var numsSize1, inputNum1 int
		fmt.Scan(&numsSize1)
		var numsList1 []int
		for j := 0; j < numsSize1; j++ {
			fmt.Scan(&inputNum1)
			numsList1 = append(numsList1, inputNum1)
		}
		var numsCount1 int
		fmt.Scan(&numsCount1)
		var numsSize2, inputNum2 int
		fmt.Scan(&numsSize2)
		var numsList2 []int
		for k := 0; k < numsSize2; k++ {
			fmt.Scan(&inputNum2)
			numsList2 = append(numsList2, inputNum2)
		}
		var numsCount2 int
		fmt.Scan(&numsCount2)
		merge(numsList1, numsCount1, numsList2, numsCount2)
		fmt.Println(numsList1)
	}
}
