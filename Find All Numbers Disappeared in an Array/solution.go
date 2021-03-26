func findDisappearedNumbers(nums []int) []int {
    numsNew := []int{}
    res := 0
    for i := 0; i < len(nums); i++ {
        res = changeInt(nums[i]) - 1
        nums[res] = changeInt(nums[res]) * -1
    }
    fmt.Println(nums)
    for j := 0; j < len(nums); j++ {
        if nums[j] > 0 {
            numsNew = append(numsNew, j+1)
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