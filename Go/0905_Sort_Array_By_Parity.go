func sortArrayByParity(A []int) []int {
    var left int = 0 
    right := len(A) - 1
    for left < right {
        if A[left] % 2 == 1 {
            if A[right] % 2 == 0 {
                A[left], A[right] = A[right], A[left]
            } else {
                right--
            }
        } else {
            left ++
        }
    }
    return A
}
