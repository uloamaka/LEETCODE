func findKDistantIndices(nums []int, key int, k int) []int {
    ans := []int{}
    n := len(nums)
    m := make(map[int]int)
    for idx, num := range nums {
        if num == key {
            for i := idx; i < n; i++ {
                l := max(0, i-k)
                r := min(n-1, i+k)
                if math.Abs(float64(l - idx)) <= float64(k) {
                    if _, ok := m[l]; !ok {
                        m[l] = l
                    }
                }
                if math.Abs(float64(r - idx)) <= float64(k) {
                    if _, ok := m[r]; !ok {
                        m[r] = r
                    }
                }
            }
        }
    }
    for _, val := range m {
        ans = append(ans, val)
    }
    sort.Ints(ans)
    return ans
}
