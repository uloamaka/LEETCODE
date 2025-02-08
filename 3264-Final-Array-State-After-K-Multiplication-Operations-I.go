func getFinalState(nums []int, k int, multiplier int) []int {
    for i := 0; i < k; i ++ {
        index, x := minIntSlice(nums)
        nx := x * multiplier
        nums[index] = nx
    }
    return nums
}


func minIntSlice(v []int) (minIndex, minValue int) {
    if len(v) == 0 {
        return 0, 0
    }
    
    minValue = v[0]
    minIndex = 0
    
    for i, e := range v {
        if e < minValue {
            minValue = e
            minIndex = i
        }
    }
    return minIndex, minValue
}