func subarraysWithKDistinct(nums []int, k int) int {
    return atMost(nums, k) - atMost(nums, k-1)
}

func atMost(nums []int, k int) int {
    freq := make(map[int]int)

    left := 0
    totalCnt := 0

    for right := 0; right < len(nums); right++ {
        freq[nums[right]]++

        for len(freq) > k {
            freq[nums[left]]--

            if freq[nums[left]] == 0 {
                delete(freq, nums[left])
            }

            left++
        }

        totalCnt += right - left + 1
    }

    return totalCnt
}