func findMaxConsecutiveOnes(nums []int) int {
    count := 0 
    maxCount := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            count = count + 1
            if count > maxCount {
                maxCount = count
            }
        } else if nums[i] == 0 {
            count = 0
        }
    }
    return maxCount
}