func canWinNim(n int) bool {
    if n < 4 {
        return true
    } else if n % 4 > 0 {
        return true
    } else {
        return false
    }
}
