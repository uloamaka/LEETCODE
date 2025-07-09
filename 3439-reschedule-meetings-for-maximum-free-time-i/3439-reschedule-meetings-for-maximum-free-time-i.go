func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	n := len(startTime)
	res := 0
	t := 0
	for i := 0; i < n; i++ {
		t += endTime[i] - startTime[i]
		var left int
		if i <= k-1 {
			left = 0
		} else {
			left = endTime[i-k]
		}
		var right int
		if i == n-1 {
			right = eventTime
		} else {
			right = startTime[i+1]
		}
		if right-left-t > res {
			res = right - left - t
		}
		if i >= k-1 {
			t -= endTime[i-k+1] - startTime[i-k+1]
		}
	}
	return res
}