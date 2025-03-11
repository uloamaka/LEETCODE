func numberOfAlternatingGroups(colors []int) int {
    n := len(colors)
    cnt := 0
    prev, curr, next := colors[n-1], colors[0], colors[1]

    for i := 0; i < n; i++ {
        if curr != prev && curr != next {
            cnt++
        }
        prev, curr, next = curr, next, colors[(i+2)%n] 
    }

    return cnt
}