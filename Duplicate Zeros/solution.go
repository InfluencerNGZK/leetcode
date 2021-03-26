func duplicateZeros(arr []int)  {
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