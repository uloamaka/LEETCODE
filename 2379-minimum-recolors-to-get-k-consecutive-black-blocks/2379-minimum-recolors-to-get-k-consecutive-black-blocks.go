func minimumRecolors(blocks string, k int) int {
    blockCnt := len(blocks)
    minCnt := k
    numIteration := 0
    if blockCnt == k {
        numIteration = 1
    } else {
        numIteration = blockCnt - k + 1
    }
    
    for n := 0; n < numIteration; n++ {
        cnt := 0 
        for _, oc := range blocks[n:n+k] {
            if oc == 'W' {
                cnt++
            }
        }
        fmt.Println(cnt)
        if cnt < minCnt {
            minCnt = cnt
        }
    }
    return minCnt
}