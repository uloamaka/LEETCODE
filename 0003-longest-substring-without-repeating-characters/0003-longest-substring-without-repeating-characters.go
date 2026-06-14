func lengthOfLongestSubstring(s string) int {
    var maxLen int
    var freq[256]int
    var l int

    for r := 0; r<len(s); r++ {
        freq[s[r]]++
       
        for freq[s[r]] > 1 {
            freq[s[l]]--
            l++
        }

        if r-l+1 > maxLen {
            maxLen = r-l+1
        }
    }

    return maxLen
}
