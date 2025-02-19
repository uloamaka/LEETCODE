func minOperations(nums []int, k int) int {
    sort.Ints(nums)
    for i, val := range nums {
        if val >= k {
            return i
        }
    }
    return 0
}
