func moveZeroes(nums []int)  {
    l, r := 0, 1
    for r < len(nums) {
        if nums[l] != 0 {
            l++
        }
        if nums[l] == 0 && nums[r] != 0 {
            nums[l], nums[r] = nums[r], nums[l]
        }  
        r++
    }
}