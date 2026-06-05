func containsDuplicate(nums []int) bool {
    if len(nums) == 0 {
        return false
    }

    m := map[int]bool{}
    for _, num := range nums{
        if _, ok := m[num]; !ok {
            m[num] = true
            continue
        }
        return true
    }
    return false
}