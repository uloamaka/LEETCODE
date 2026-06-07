
func maximumSubarraySum(nums []int, k int) int64 {
    n := len(nums)
    if k > n {
        return 0
    }

    freq := make(map[int]int)
    windowSum := int64(0)
    maxSum := int64(0)

    for i := 0; i < n; i++ {
        freq[nums[i]]++
        windowSum += int64(nums[i])

        if i >= k {
            left := nums[i-k]
            freq[left]--
            if freq[left] == 0 {
                delete(freq, left)
            }
            windowSum -= int64(left)
        }

        if i >= k-1 && len(freq) == k {
            if windowSum > maxSum {
                maxSum = windowSum
            }
        }
    }

    return maxSum
}