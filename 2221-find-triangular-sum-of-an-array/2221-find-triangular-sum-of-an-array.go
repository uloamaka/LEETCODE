func triangularSum(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }

    for i := 0; i < n-1; i++{
        nums[i] = (nums[i] + nums[i+1]) % 10
    }

    return triangularSum(nums[:n-1]) 
}