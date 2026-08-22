func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    subset := []int{}
    return subsetFinder(nums, 0, subset)
}

func subsetFinder(nums []int, idx int, subset []int) [][]int {
    n := len(nums)
    ans := [][]int{}

    if idx == n {
        copySubset := append([]int{}, subset...)
        ans = append(ans, copySubset)
        return ans
    }

    subset = append(subset, nums[idx])
    ans = append(ans, subsetFinder(nums, idx+1, subset)...)

    subset = subset[:len(subset)-1]
    for idx < n-1 && nums[idx] == nums[idx+1] {
        idx++
    }
    ans = append(ans, subsetFinder(nums, idx+1, subset)...)

    return ans
}
/////////////////////////////////////////////////////////
/* func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    var queue [][]int
    queue = append(queue, []int{})
    
    m := make(map[int]int)
    
    for i := 0; i < len(nums); i++ {
        length := len(queue)
        
        limit := -1
        
        if val, ok := m[nums[i]]; ok {
            limit = val
        } else {
            m[nums[i]] = length
        }
        
        m[nums[i]] = length
        
        for j := 0; j < length; j++ {
            if limit != -1 && j < limit {
                continue    
            }
            
            item := queue[j]

            tmp := make([]int, len(item))
            copy(tmp, item)
            tmp = append(tmp, nums[i])

            queue = append(queue, tmp) 
        }
    }
    
    
    return queue
} */