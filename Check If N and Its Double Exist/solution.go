func checkIfExist(arr []int) bool {
    a := 0
    for i := 0; i < len(arr); i++ {
        a = arr[i]
        for j := i+1; j < len(arr); j++ {
            if arr[j] *2 == a || a * 2 == arr[j] {
                return true
            }
        }
    }
    return false
}