func subsetXORSum(nums []int) int {
	sumTotal := 0

	for _, num := range nums {
		sumTotal |= num
	}
	return sumTotal << (len(nums) - 1)
}
