func minSubArrayLen(target int, nums []int) int {
    minLen := math.MaxInt32
    winSum := 0

    l := 0
    for r := 0; r < len(nums); r++ {
        winSum += nums[r]

        for winSum >= target {
            curLen := r - l + 1
            if curLen < minLen {
                minLen = curLen
            }
           
            winSum -= nums[l]
            l++
        }
    }

    if minLen == math.MaxInt32 {
        minLen = 0
    }
    return minLen
}