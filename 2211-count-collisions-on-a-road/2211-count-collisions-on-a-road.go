func countCollisions(directions string) int {
    n := len(directions)
    l, r := 0, n-1
    for l < n {
        if directions[l] != 'L' {
            break
        }
        l++
    }
    for r >= 0 {
        if directions[r] != 'R' {
            break
        }
        r--
    }
    total := 0
    for i := l; i <= r; i++ {
        if directions[i] != 'S' {
            total++
        }
    }
    return total
}