func countValidSelections(nums []int) int {
    cnt, nonZeros, n := 0, 0, len(nums)
    for _, num := range nums {
        if num > 0 {
            nonZeros++
        }
    }
    for i :=0; i < n; i++ {
        if nums[i] == 0 {
            if isValid(nums, nonZeros, i, -1) {
                cnt++
            }
            if isValid(nums, nonZeros, i, 1) {
                cnt++
            }
        }
    }
    return cnt
}

func isValid(nums []int, nonZeros, start, direction int) bool {
    temp := make([]int, len(nums))
    copy(temp, nums)
    curr := start
    for nonZeros > 0 && curr >= 0 && curr < len(nums) {
        if temp[curr] > 0 {
            temp[curr]--
            direction *= -1
            if temp[curr] == 0 {
                nonZeros--
            }
        }
        curr += direction
    }
    return nonZeros == 0
}