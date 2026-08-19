func runningSum(nums []int) []int {
    ans := []int{}
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += recur(nums, i)
        ans = append(ans, sum)
    }
   
    return ans
}

func recur(nums []int, n int) int {
    if n == len(nums) {
        return 0
    }
    return nums[n]
}