func maxProductDifference(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    return (nums[n-2] * nums[n - 1]) - (nums[1] * nums[0]) 
}