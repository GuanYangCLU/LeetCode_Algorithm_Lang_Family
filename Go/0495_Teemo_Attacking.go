func findPoisonedDuration(timeSeries []int, duration int) int {
    if (len(timeSeries) == 0) {
        return 0
    }
    var t int = len(timeSeries) * duration
    for i:=0;i < len(timeSeries) - 1;i++ {
        if (timeSeries[i] + duration > timeSeries[i+1]) {
            t -= (timeSeries[i] + duration - timeSeries[i+1])
        }
    }
    return t
}
