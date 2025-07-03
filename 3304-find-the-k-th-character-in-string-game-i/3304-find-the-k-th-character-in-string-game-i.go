func kthCharacter(k int) byte {
    if k == 1 {
		return 'a'
	}

	// Find the largest power of 2 less than k
	lengthPrev := 1
	for lengthPrev*2 < k {
		lengthPrev *= 2
	}

	// The k-th character is the next character of the (k-lengthPrev)-th character in the previous step
	prevChar := kthCharacter(k - lengthPrev)
	return 'a' + (prevChar-'a'+1)%26
}