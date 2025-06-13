func minimizeMax(nums []int, p int) int {
    sort.Ints(nums)

    left := 0
    right := nums[len(nums)-1] - nums[0]
    result := right
    for left <= right {
        mid := left + (right - left) / 2
        if countValidPairs(nums, mid) >= p {
            result = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return result
}

func countValidPairs(nums []int, mid int) int {
    count := 0
	i := 0
	for i < len(nums)-1 {
		if nums[i+1]-nums[i] <= mid {
			count++
			i += 2 
		} else {
			i++
		}
	}
	return count
}