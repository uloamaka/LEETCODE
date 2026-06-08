func maximumUniqueSubarray(nums []int) int {
    var max, l int

    seen := [10001]bool{}
    var summer int
    for r := 0; r < len(nums); r++ {
        for seen[nums[r]] {
            seen[nums[l]] = false
            summer -= nums[l]
            l++
        }

        seen[nums[r]] = true
        summer += nums[r]

        if summer > max {
            max = summer
        } 
    } 
    return max
}
    