func minSubArrayLen(target int, nums []int) int {
    minLen := len(nums)
    summer, cnt := 0, 0

    l := 0
    for r := 0; r < len(nums); r++ {
        cnt++
        summer += nums[r]

        for summer >= target {
            //check and update minLen
            if cnt < minLen {
                minLen = cnt
            }
            // remove the value of the leaving left pnter from summer
            summer -= nums[l]
            // move left pnter
            l++
            // decreament cnter
            cnt--
        }
        
        if summer < target && cnt == len(nums) {
            minLen = 0
        }
    }

    return minLen
}