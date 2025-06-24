func findKDistantIndices(nums []int, key int, k int) []int {
    ans := []int{}
    n := len(nums)
    m := make(map[int]bool)
    for j, num := range nums {
        if num == key {
            l := max(0, j-k)
            r := min(n-1, j+k)
            for i := l; i <= r; i++ {
				m[i] = true
			}
    
        }
    }
    
    for val := range m {
        ans = append(ans, val)
    }

    sort.Ints(ans)
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
