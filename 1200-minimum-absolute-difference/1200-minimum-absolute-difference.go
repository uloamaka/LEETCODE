func minimumAbsDifference(arr []int) [][]int {
    if sort.IntsAreSorted(arr) != true {
        sort.Ints(arr)
    } 
    minDiff := math.MaxInt32
    var ans [][]int
    
    for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		
		if diff < minDiff {
			minDiff = diff
			ans = [][]int{{arr[i-1], arr[i]}} // Reset with new pair
		} else if diff == minDiff {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}
    return ans 
}


