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
        return count+1
    } else {
        return 0
    }
}