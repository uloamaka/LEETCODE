func smallestNumber(n int) int {
    g := 1
    for g < n {
        g = g*2 + 1
    }
    return g
}
