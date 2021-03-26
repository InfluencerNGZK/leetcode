func findNumbers(nums []int) int {
    count := 0
    for i := range(nums) {
        if (nums[i] / 10) != 0 {
            if (nums[i] / 100) == 0 {
                count = count + 1
            } else if (nums[i] / 1000) != 0 {
                if (nums[i] / 10000) == 0 {
                    count = count + 1
                } else if (nums[i] / 100000) != 0 {
                    if (nums[i] / 1000000) == 0 {
                        count = count + 1
                    }
                }
            }
        }
    }
    return count
}