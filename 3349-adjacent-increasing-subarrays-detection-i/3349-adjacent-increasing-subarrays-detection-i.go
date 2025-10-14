func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)
	cnt, precnt, ans := 1, 0, 0
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			cnt++
		} else {
			precnt = cnt
			cnt = 1
		}
		ans = max(ans, min(precnt, cnt))
		ans = max(ans, cnt/2)
	}
	return ans >= k
}