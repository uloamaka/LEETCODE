func runningSum(nums []int) []int {
    ans := []int{}

    for i := 0; i < len(nums); i++ {
        ans = append(ans, recur(nums, i))
    }
   
    return ans
}

func recur(nums []int, n int) int {
    if n == 0 {
        return nums[0]
    }
    return nums[n] + recur(nums, n-1)
}
