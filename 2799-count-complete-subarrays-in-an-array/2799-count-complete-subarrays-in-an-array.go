func countCompleteSubarrays(nums []int) int {
    freq := make(map[int]int)
    cnt := 0
    for _, num := range nums {
        freq[num]++
    }

    seen := make(map[int]int) 
    left := 0
    for right := 0; right < len(nums); right++ {
        seen[nums[right]]++
        for len(seen) == len(freq) {
            seen[nums[left]]--

            if seen[nums[left]] == 0 {
                delete(seen, nums[left])
            }

            left++
            cnt += len(nums)-right
        }
    }

    return cnt
}