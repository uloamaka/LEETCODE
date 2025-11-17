func kLengthApart(nums []int, k int) bool {
    last := -1 // index of last seen 1

    for i, num := range nums {
        if num == 1 {
            if last != -1 && i-last-1 < k {
                return false
            }
            last = i
        }
    }
    return true
}