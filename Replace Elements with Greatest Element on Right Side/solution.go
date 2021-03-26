func replaceElements(arr []int) []int {
    arrNew := make([]int, len(arr))
    for i := 0; i < len(arr)-1; i++ {
        a := 0
        a = arr[i+1]
        for j := i+1; j < len(arr); j++ {
            if arr[j] > a {
                a = arr[j]
            }
        }
        arrNew[i] = a
    }
    arrNew[len(arr)-1] = -1
    return arrNew
}