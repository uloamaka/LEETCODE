func longestSubsequence(s string, k int) int {
    // n := len(s)
    
    // Count total zeros - they can always be included
    zeros := 0
    for _, ch := range s {
        if ch == '0' {
            zeros++
        }
    }
    
    // Try different numbers of ones, starting from maximum possible
    maxOnes := 0
    for numOnes := min(32, countOnes(s)); numOnes >= 0; numOnes-- {
        if canFormValue(s, numOnes, k) {
            maxOnes = numOnes
            break
        }
    }
    
    return zeros + maxOnes
}

func countOnes(s string) int {
    count := 0
    for _, ch := range s {
        if ch == '1' {
            count++
        }
    }
    return count
}

func canFormValue(s string, numOnes int, k int) bool {
    if numOnes == 0 {
        return true
    }
    
    // Greedy: pick the rightmost 'numOnes' ones to minimize value
    value := 0
    power := 1
    onesUsed := 0
    
    // Scan from right to left
    for i := len(s) - 1; i >= 0 && onesUsed < numOnes; i-- {
        if s[i] == '1' {
            value += power
            onesUsed++
            if value > k {
                return false
            }
        }
        power *= 2
        if power > k + 1 { // Prevent overflow
            break
        }
    }
    
    return onesUsed == numOnes && value <= k
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}