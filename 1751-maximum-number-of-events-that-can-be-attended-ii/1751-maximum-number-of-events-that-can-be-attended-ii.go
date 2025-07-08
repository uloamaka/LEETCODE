func maxValue(events [][]int, k int) int {
    n := len(events)
    if n == 0 || k == 0 {
        return 0
    }
    
    sort.Slice(events, func(i, j int) bool {
        return events[i][0] < events[j][0]
    })
    
    dp := make([][]int, k+1)
    for i := range dp {
        dp[i] = make([]int, n)
        for j := range dp[i] {
            dp[i][j] = -1 
        }
    }
    
    nextIndices := make([]int, n)
    for i := 0; i < n; i++ {
        nextIndices[i] = findNextIndex(events, i)
    }
    
    var dfs func(curIndex, count int) int
    dfs = func(curIndex, count int) int {
        if count == 0 || curIndex == n {
            return 0
        }
        
        
        if dp[count][curIndex] != -1 {
            return dp[count][curIndex]
        }
        
        // Option 1: Skip this event and get value of dfs(curIndex + 1, count)
        skipValue := dfs(curIndex+1, count)
        
        // Option 2: Attend this event
        // Get the index of nearest available event after current event
        nextIndex := nextIndices[curIndex]
        // Get value of dfs(nextIndex, count - 1) plus value of this event
        attendValue := events[curIndex][2] + dfs(nextIndex, count-1)
        
        // Assign the larger value between the two options
        if skipValue > attendValue {
            dp[count][curIndex] = skipValue
        } else {
            dp[count][curIndex] = attendValue
        }
        
        return dp[count][curIndex]
    }
    
    return dfs(0, k)
}

// Helper function to find the next available event index
func findNextIndex(events [][]int, curIndex int) int {
    n := len(events)
    currentEndTime := events[curIndex][1]
    
    // Binary search for the first event that starts after current event ends
    left, right := curIndex+1, n
    
    for left < right {
        mid := left + (right-left)/2
        if events[mid][0] > currentEndTime {
            right = mid
        } else {
            left = mid + 1
        }
    }
    
    return left
}