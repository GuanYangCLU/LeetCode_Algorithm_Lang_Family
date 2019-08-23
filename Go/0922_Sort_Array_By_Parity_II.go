func sortArrayByParityII(A []int) []int {
    rs := []int{}
    if len(A) == 0 {
        return rs
    }
    var even [] int
    var odd [] int
    for i:= 0;i<len(A);i++ {
        if A[i] % 2 == 1 {
            odd = append(odd, 0, A[i])
        } else {
            even = append(even, A[i], 0)
        }
    }
    for j:=0;j<len(A);j++ {
        rs = append(rs, odd[j] + even[j])
    }
    return rs
}
