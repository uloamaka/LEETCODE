func divideArray(nums []int, k int) [][]int {
    sort.Ints(nums)
    
    n := len(nums)
    ans := [][]int{}
    left, right := 0, 3
    for i := 0; i < n; i += 3 {
        if nums[right-1] - nums[left] <= k {
            ans = append(ans, nums[left:right])
            left = right
            right = right + 3
        } else {
            return [][]int{}
        }
    }
    return ans
}

