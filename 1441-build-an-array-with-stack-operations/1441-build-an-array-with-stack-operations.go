func buildArray(target []int, n int) []string {
    ops := []string{}
 
    j := 0
    for i := 1; i <= n; i++ {
        if j == len(target) {
            break
        }
        ops = append(ops, "Push")
        if i == target[j] {
            j++
        } else  {
            ops = append(ops, "Pop")
        }
    }
    return ops
}