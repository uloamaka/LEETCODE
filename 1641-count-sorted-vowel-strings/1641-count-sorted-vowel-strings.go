func countVowelStrings(n int) int {
    return comb(n+4, 4)
}

func comb(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	// efficient iterative version
	res := 1
	for i := 1; i <= k; i++ {
		res = res * (n - i + 1) / i
	}
	return res
}
