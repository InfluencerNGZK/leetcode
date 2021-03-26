func merge(nums1 []int, m int, nums2 []int, n int)  {
    for i := m; i < m+n; i++ {
        nums1[i] = nums2[i-m]
    }
    for k := range(nums1) {
		for j := 1; j < len(nums1)-k; j++ {
			if nums1[j] < nums1[j-1] {
                a := nums1[j]
                nums1[j] = nums1[j-1]
                nums1[j-1] = a
			}
		}
	}
}