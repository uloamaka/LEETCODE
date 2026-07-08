func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 || k == 0 {
		return nil
	}

    result := make([]int, 0, len(nums)-k+1)
    deque := make([]int, 0)

    for i :=0; i<len(nums); i++ {
        // Remove indices that are outside the current window
        if len(deque) > 0 && deque[0] <= i-k {
            deque = deque[1:]
        }

        // Remove all smaller (or equal) values from the back
		// because they can never become the maximum
        for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
            deque = deque[:len(deque)-1]
        }

        // add the current index
        deque = append(deque, i)

        // Once we've formed the first full window,
		// record the maximum (front of deque)
        if i >= k-1 {
            result = append(result, nums[deque[0]])
        }
    }
    return result
}
