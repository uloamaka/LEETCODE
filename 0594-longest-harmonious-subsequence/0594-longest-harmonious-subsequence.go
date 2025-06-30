func findLHS(nums []int) int {
    freq := make(map[int]int)
    for _, num := range nums {
        freq[num]++
    }
    
    maxLen := 0
    for key, count := range freq {
        if countAdj, exists := freq[key + 1]; exists {
            if count + countAdj > maxLen {
                maxLen = count + countAdj
            }
        }
    }
    return maxLen
}