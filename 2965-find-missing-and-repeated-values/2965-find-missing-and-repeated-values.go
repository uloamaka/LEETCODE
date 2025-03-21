func findMissingAndRepeatedValues(grid [][]int) []int {
    sum, sqrSum := 0, 0
    n := len(grid)
    total := n*n
    for r := 0; r < n; r++ {
        for c := 0; c< n; c++ {
            sum += grid[r][c]
            sqrSum += grid[r][c]*grid[r][c]
        }
    }
    sumDiff := sum - (total * (total + 1) / 2)
    sqrDiff := sqrSum - (total * (total + 1) * (2 * total + 1) / 6)
    
    repeat := (sqrDiff / sumDiff + sumDiff) / 2
    missing := (sqrDiff / sumDiff - sumDiff) / 2
    
    return []int{repeat, missing}
}