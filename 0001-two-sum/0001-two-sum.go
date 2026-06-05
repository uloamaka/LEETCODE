func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    slc := []int{}
    for i, num := range nums {
        diff := target-num
        if j, seen := m[diff]; seen {
            slc = append(slc, i, j)
            break
        }
        m[num] = i
    }
    return slc
}