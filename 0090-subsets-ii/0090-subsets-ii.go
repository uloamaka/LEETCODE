func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    ans := [][]int{}
    subset := []int{}

    var dfs func(i int, subset []int)

    dfs = func(i int, subset []int) {
        n := len(nums)
        if i == n {
            copySubset := append([]int{}, subset...)
            ans = append(ans, copySubset)
            return 
        }

        subset = append(subset, nums[i])
        dfs(i+1, subset)

        subset = subset[:len(subset)-1]
        for i < n-1 && nums[i] == nums[i+1] {
            i++
        }
        dfs(i+1, subset)
    }

    dfs(0, subset)

    return ans 
}
