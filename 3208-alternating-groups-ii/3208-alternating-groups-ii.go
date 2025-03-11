func numberOfAlternatingGroups(colors []int, k int) int {
    n := len(colors)
    mismatches := 0
    for i := 0; i < k-1; i++ {
        if colors[i] != colors[i+1] {
            mismatches++
        }
    }
    
    cnt := 0
    if mismatches == k-1 {
        cnt++
    }

    for i := 1; i < n; i++ {
        if colors[i-1] != colors[i] {
            mismatches--
        }
        if colors[(i+k-2) % n] != colors[(i+k-1) % n] {
            mismatches++
        }
        if mismatches == k-1 {
            cnt++
        }
    }

    return cnt
}