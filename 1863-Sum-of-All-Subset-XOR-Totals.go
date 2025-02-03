func subsetXORSum(nums []int) int {
    return dfs(nums, 0, 0)
}

func dfs(nums []int, index int, xorSum int) int {
    if index == len(nums) {
        return xorSum 
    }
    
    include := dfs(nums, index+1, xorSum^nums[index])
    
    exclude := dfs(nums, index+1, xorSum)
    
    return include + exclude
}