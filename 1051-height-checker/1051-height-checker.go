func heightChecker(heights []int) int {
    cnt := 0
    srtHeights := append([]int(nil), heights...)
    sort.Ints(srtHeights)
    for i := 0; i < len(heights); i++ {
        if heights[i] != srtHeights[i] {
            cnt++
        }
    }
    return cnt
}