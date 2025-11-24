func prefixesDivBy5(nums []int) []bool {
	ans := make([]bool, len(nums))
	remainder := 0

	for i, num := range nums {
		remainder = remainder * 2
		remainder = remainder + num
		remainder = remainder % 5

		ans[i] = remainder == 0
	}

	return ans
}