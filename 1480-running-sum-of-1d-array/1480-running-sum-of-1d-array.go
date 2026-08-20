// func runningSum(nums []int) []int {
//     ans := []int{}

//     for i := 0; i < len(nums); i++ {
//         ans = append(ans, recur(nums, i))
//     }
   
//     return ans
// }

// func recur(nums []int, n int) int {
//     if n == 0 {
//         return nums[0]
//     }
//     return nums[n] + recur(nums, n-1)
// }

func recur(nums []int, n int, sum int, ans []int) []int {
    if n == len(nums) {
        return ans
    }

    sum += nums[n]
    ans = append(ans, sum)

    return recur(nums, n+1, sum, ans)
}

func runningSum(nums []int) []int {
    return recur(nums, 0, 0, []int{})
}
