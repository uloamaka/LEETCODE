func lengthOfLongestSubstring(s string) int {
    var maxLen int
    var freq[256]int
    var l int

    for r := 0; r<len(s); r++ {
        freq[s[r] - 'a']++
       
        for freq[s[r] - 'a'] > 1 {
            freq[s[l]-'a']--
            l++
        }

        if r-l+1 > maxLen {
            maxLen = r-l+1
        }
    }

    return maxLen
}
