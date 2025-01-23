func buildArray(nums []int) []int {
    arrLen := len(nums)
    ans := make([]int, arrLen)
    for i, num := range nums {
        ans[i] = nums[num]
    }

    return ans
}
// algorithm---------------------
// Create a new array of same length as input
// For each index, use the element's value as a secondary index
// Retrieve value from original array using secondary indexing
// Populate new array with these retrieved values