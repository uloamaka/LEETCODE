func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i, num := range nums {
        diff := target-num
        if j, seen := m[diff]; seen {
            return []int{i, j}
        }
        m[num] = i
    }
    return []int{}
}