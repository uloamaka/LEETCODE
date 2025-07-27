func countHillValley(nums []int) int {
    cnt := 0
    NonEqualNums := []int{nums[0]}
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] {
            NonEqualNums = append(NonEqualNums, nums[i])
        } 
    }

    n := len(NonEqualNums)

    for i := 1; i < n -1; i++ { 
        if NonEqualNums[i] > NonEqualNums[i-1] && NonEqualNums[i] > NonEqualNums[i+1] {
            cnt++
        }
        if NonEqualNums[i] < NonEqualNums[i-1] && NonEqualNums[i] < NonEqualNums[i+1] {
            cnt++
        }
    }

    return cnt
}