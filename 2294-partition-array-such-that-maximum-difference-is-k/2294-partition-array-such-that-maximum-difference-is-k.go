func partitionArray(nums []int, k int) int {
    cnt := 1
    start := nums[0]
    
    for i := 1; i < len(nums); i++ {
        if nums[i]-start > k {
            cnt++
            start = nums[i]
        }
    }
    return cnt
}