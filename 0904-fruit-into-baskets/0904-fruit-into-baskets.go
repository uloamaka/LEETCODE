func totalFruit(fruits []int) int {
	var l, p, ans int
	for i := 1; i < len(fruits); i++ {
		if fruits[i] != fruits[i-1] {
            if p > 0 && fruits[i] != fruits[p-1] {
                l = p
            }
			p = i
		}
        ans = max(ans, i - l)
	}
	return ans + 1
}