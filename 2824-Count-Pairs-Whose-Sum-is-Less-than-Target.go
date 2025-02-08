func countPairs(nums []int, target int) int {
    count := 0
    left := 0
    right := len(nums) - 1
    sort.Ints(nums)
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