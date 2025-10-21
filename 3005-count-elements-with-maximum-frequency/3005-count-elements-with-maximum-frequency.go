func maxFrequencyElements(nums []int) int {
    freqM := make(map[int]int)
    for _, num := range nums {
        freqM[num]++
    }
    maxF, cnt := 0, 0 
    for _, cn := range freqM {
        if cn > maxF {
            maxF = cn
            cnt = 1
        } else if cn == maxF {
            cnt++
        }
    }
    return maxF*cnt
}