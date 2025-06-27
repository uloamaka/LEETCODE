func longestSubsequenceRepeatedK(s string, k int) string {
    n := len(s)
	freq := make(map[byte]int)
	for i := 0; i < n; i++ {
		freq[s[i]]++
	}

	// Reduce character set: only chars appearing at least k times
	var chars []byte
	for c, f := range freq {
		if f >= k {
			chars = append(chars, c)
		}
	}

	// Sort descending for lexicographical preference
	sort.Slice(chars, func(i, j int) bool { return chars[i] > chars[j] })

	// Helper function: check if t*k is a subsequence of s
	isValid := func(t string) bool {
		target := ""
		for i := 0; i < k; i++ {
			target += t
		}

		idx := 0
		for i := 0; i < len(s) && idx < len(target); i++ {
			if s[i] == target[idx] {
				idx++
			}
		}
		return idx == len(target)
	}

	queue := []string{""}
	var res string

	// BFS Loop
	for len(queue) > 0 {
		nextQueue := []string{}
		for _, cur := range queue {
			for _, c := range chars {
				candidate := cur + string(c)
				if isValid(candidate) {
					// Update result if it's longer or lexicographically larger
					if len(candidate) > len(res) || (len(candidate) == len(res) && candidate > res) {
						res = candidate
					}
					// Only expand further if it's not too long
					if len(candidate) <= n/k {
						nextQueue = append(nextQueue, candidate)
					}
				}
			}
		}
		queue = nextQueue
	}

	return res
}