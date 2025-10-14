func hasIncreasingSubarrays(nums []int, k int) bool {
    n := len(nums)
    l := 0
    ans1, ans2 := false, false
    for l <= n-2*k {
        r := l+k

        fseen1 := -1000
        for _, val := range nums[l:l+k] {
            if fseen1 >= val {
                ans1 = false
                break
            } else {
                fseen1 = val
                ans1 = true
            }
        }
        
        fseen2 := -1000
        for _, val := range nums[r:r+k] {
            if fseen2 >= val {
                ans2 = false
                break
            } else {
                fseen2 = val
                ans2 = true
            }
        }

        if ans1 == true && ans2 == true {
            return true
        }

        l++
    }
    return false
}