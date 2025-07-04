func kthCharacter(k int64, operations []int) byte {
	if k == 1 {
		return 'a'
	}

	// Find the largest power of 2 less than k
	var lengthPrev int64 = 1
	for lengthPrev*2 < k {
		lengthPrev *= 2
	}

	// Determine which operation caused the doubling from lengthPrev to lengthPrev*2
	opIndex := 0
	temp := lengthPrev
	for temp > 1 {
		temp >>= 1
		opIndex++
	}

	// Handle case where operations might be shorter than expected
	if opIndex >= len(operations) {
		opIndex = len(operations) - 1
	}

	prevChar := kthCharacter(k-lengthPrev, operations)
	if operations[opIndex] == 0 {
		return prevChar
	}
	// Operation 1: increment character with wrap-around
	return 'a' + (prevChar-'a'+1)%26
}