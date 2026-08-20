func subsets(nums []int) [][]int {
    ans := [][]int{}
    subset := []int{}

    var dfs func(n int)

    dfs = func(n int) {
        if n == len(nums) {
            copySubset := append([]int{}, subset...)
            ans = append(ans, copySubset)
            return
        }

        // Include nums[n]
        subset = append(subset, nums[n])
        dfs(n + 1)

        // Don't include nums[n]
        subset = subset[:len(subset)-1]
        dfs(n + 1)
    }

    dfs(0)
    return ans
}