func maxSubsequence(nums []int, k int) []int {
    n := len(nums)
	vals := make([][]int, n) // auxiliary array
	for i := 0; i < n; i++ {
		vals[i] = []int{i, nums[i]}
	}
	// sort by numerical value in descending order
	sort.Slice(vals, func(i, j int) bool {
		return vals[i][1] > vals[j][1]
	})
	// select the first k elements and sort them in ascending order by index
	sort.Slice(vals[:k], func(i, j int) bool {
		return vals[i][0] < vals[j][0]
	})
	res := make([]int, k) // target subsequence
	for i := 0; i < k; i++ {
		res[i] = vals[i][1]
	}
	return res
}
// sort the array while maintainting their index 
// using a hash map
// then iterate through the map getting their highest keys that sums up to maxSub 
// then return it 