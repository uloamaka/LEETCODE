func containsDuplicate(nums []int) bool {
    slices.Sort(nums)
    
    count := len(nums)
    seen := nums[0]

    for i := 1; i < count; i++ {
       if seen == nums[i] {
            return true
        }
        seen = nums[i]
    }  
    return false
}