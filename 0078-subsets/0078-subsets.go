func subsets(nums []int) [][]int {
    ans, subset := [][]int{}, []int{}

    var dfs func(n int)

    dfs = func(n int) {
        if n >= len(nums){
            ans = append(ans, append([]int{}, subset...))
            return 
        }

        subset = append(subset, nums[n])
        dfs(n + 1)

        subset = subset[:len(subset)-1]
        dfs(n + 1)
    }

    dfs(0)
    return ans 
}