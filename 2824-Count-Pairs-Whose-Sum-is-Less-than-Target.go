func countPairs(nums []int, target int) int {
    count := 0
    sort.Ints(nums)
    left := 0
    right := len(nums) - 1
    for left < right {  
        if nums[right] + nums[left] < target {
            count += right - left
            left++
        } else {
            right--
        }
    }
    return count
}