func maximumSubarraySum(nums []int, k int) int64 {
	var maxSum int64 = 0
	var currentSum int64 = 0
	left := 0

	seen := make(map[int]bool)

	for right := 0; right < len(nums); right++ {
		for seen[nums[right]] || (right-left+1 > k) {
			delete(seen, nums[left])
			currentSum -= int64(nums[left])
			left++
		}

		seen[nums[right]] = true
		currentSum += int64(nums[right])

		if right-left+1 == k {
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}

	return maxSum
}
