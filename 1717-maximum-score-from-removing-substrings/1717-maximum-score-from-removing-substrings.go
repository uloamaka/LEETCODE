func maximumGain(s string, x int, y int) int {
	totalScore := 0
	highPriorityPair, lowPriorityPair := "", ""
	highValue, lowValue := 0, 0

	// Determine priority based on which value is higher
	if x > y {
		highPriorityPair, lowPriorityPair = "ab", "ba"
		highValue, lowValue = x, y
	} else {
		highPriorityPair, lowPriorityPair = "ba", "ab"
		highValue, lowValue = y, x
	}

	// First pass: remove high priority pair
	processed, removedHigh := removeSubstring(s, highPriorityPair)
	totalScore += removedHigh * highValue

	// Second pass: remove low priority pair
	_, removedLow := removeSubstring(processed, lowPriorityPair)
	totalScore += removedLow * lowValue

	return totalScore
}

func removeSubstring(input string, targetPair string) (string, int) {
	stack := []byte{}
	count := 0

	for i := 0; i < len(input); i++ {
		curr := input[i]
		if len(stack) > 0 && stack[len(stack)-1] == targetPair[0] && curr == targetPair[1] {
			stack = stack[:len(stack)-1] // pop the last character
			count++
		} else {
			stack = append(stack, curr)
		}
	}

	return string(stack), count
}