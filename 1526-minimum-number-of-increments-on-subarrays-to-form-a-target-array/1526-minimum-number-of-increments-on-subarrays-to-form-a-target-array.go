func minNumberOperations(target []int) int {
    	n := len(target)

	result := 0
	prev := 0

	for i := 0; i < n; i++ {
		if target[i] > prev {
			result += target[i] - prev
		}
		prev = target[i]
	}

	return result
}