func sortTheStudents(score [][]int, k int) [][]int {
    n := len(score)
    if n == 0 {
        return nil
    }

    indices := make([]int, n)
    for i := range indices {
        indices[i] = i
    }

    sort.Slice(indices, func(i, j int) bool {
        return score[indices[i]][k] > score[indices[j]][k]
    })

    ans := make([][]int, n)
    for i, idx := range indices {
        ans[i] = score[idx]
    }

    return ans
}