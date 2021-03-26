func moveZeroes(nums []int)  {
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