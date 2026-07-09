func subarraysWithKDistinct(nums []int, k int) int {
    fmt.Println(atMost(nums, k))
    fmt.Println(atMost(nums, k)-1)
    return atMost(nums, k) - atMost(nums, k-1)
}

func atMost(nums []int, k int) int {
    if k < 0 {
        return 0
    }

    seen := make(map[int]int)

    left := 0
    ans := 0

    for right := 0; right < len(nums); right++ {
        if seen[nums[right]] == 0 {
            k--
        }

        seen[nums[right]]++

        for k < 0 {
            seen[nums[left]]--

            if seen[nums[left]] == 0 {
                k++
            }

            left++
        }

        ans += right - left + 1
    }

    return ans
}