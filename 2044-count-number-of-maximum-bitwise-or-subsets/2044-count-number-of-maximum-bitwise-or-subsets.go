func countMaxOrSubsets(nums []int) int {
	var count int = 0
	maxOR := 0
	for _, num := range nums {
		maxOR |= num
	}
	backtrack(nums, 0, maxOR, 0, &count)
	return count
}

func backtrack(nums []int, index int, maxOR int, currentOR int, count *int) {
	if currentOR == maxOR {
		*count++
	}

	for i := index; i < len(nums); i++ {
		backtrack(nums, i+1, maxOR, currentOR|nums[i], count)
	}
}