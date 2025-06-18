func divideArray(nums []int, k int) [][]int {
    sort.Ints(nums)

    n := len(nums)
    ans := [][]int{}
    for i := 0; i < n; i += 3 {
        if nums[i+2] - nums[i] <= k {
            ans = append(ans, nums[i:i+3])
        } else {
            return [][]int{}
        }
    }
    return ans
}

