func maximumUniqueSubarray(nums []int) int {
    var max, l int

    freq := make(map[int]int)
    var summer int
    for r := 0; r < len(nums); r++ {
        freq[nums[r]]++
        summer += nums[r]

        for freq[nums[r]] > 1 {
            freq[nums[l]]--
            summer -= nums[l]
            l++
        }

        if summer > max {
            max = summer
        } 
    } 
    return max
}
    // m := make(map[int]bool)
    // ans := 0
    // for i := 0; i < len(nums); i++ {
    //     if _, ok := m[nums[i]]; !ok {
    //         m[nums[i]] = true
    //         ans += nums[i]
    //     } 
    // }
    // fmt.Println(m)
    // return ans