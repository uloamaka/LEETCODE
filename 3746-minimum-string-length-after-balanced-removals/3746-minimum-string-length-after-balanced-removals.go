func minLengthAfterRemovals(s string) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			cnt++
		}
	}
    diff := len(s) - cnt*2
	return abs(diff)
}

func abs(val int) int {
    if val < 0 {
        return -val
    }
    return val
}
