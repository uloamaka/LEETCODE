func isIncreasing(arr []int) bool {
    for i := 1; i < len(arr); i++ {
        if arr[i] <= arr[i-1] {
            return false
        }
    }
    return true
}

func hasIncreasingSubarrays(nums []int, k int) bool {
    n := len(nums)
    if n < 2*k {
        return false
    }

    for l := 0; l <= n-2*k; l++ {
        if isIncreasing(nums[l:l+k]) && isIncreasing(nums[l+k:l+2*k]) {
            return true
        }
    }
    return false
}
