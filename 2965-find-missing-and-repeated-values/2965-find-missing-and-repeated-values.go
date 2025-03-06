func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
    ans := make([]int, 2)
    dupMap := make(map[int]bool)
    for i := 0; i < n; i++ {
        for _, gdValue := range grid[i] {
            if _, ok := dupMap[gdValue]; !ok {
                dupMap[gdValue] = false
            } else {
                dupMap[gdValue] = true
                ans[0] = gdValue
            }
        }
    }
    for i := 1; i <= n*n ; i++ {
        if _, ok := dupMap[i]; !ok {
            ans[1] = i
        }
    }
    return ans
    
}