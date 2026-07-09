func firstMissingPositive(nums []int) int {
    seen := make(map[int]bool)
    max := -math.MaxInt
    for _, num := range nums {
        seen[num] = true
        if num > max {
            max = num
        }
    }

    if max < 0 {
        return 1
    }

    for i := 1; i < max; i++ {
        if _, ok := seen[i]; !ok {
            return i
        }
    } 

    return max+1
}