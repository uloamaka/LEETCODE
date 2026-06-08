func lengthOfLongestSubstring(s string) int {
    var longest, l int

    freq := make(map[byte]int)
    for r := 0; r < len(s); r++ {
        freq[s[r]]++

        for freq[s[r]] > 1 {
            freq[s[l]] -= 1
            l++
        }
        if (r - l + 1) > longest {
            longest = r-l+1
        }
    }
    return longest
}
