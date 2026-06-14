func characterReplacement(s string, k int) int {
    var freq [26]int
    var maxChar int
    var maxFreq int 
    
    var l int
    for r := 0; r < len(s); r++ {
        freq[s[r] - 'A']++

        if freq[s[r]- 'A'] > maxFreq {
            maxFreq = freq[s[r] - 'A']
        }

        if r - l + 1 - maxFreq  > k {
            freq[s[l] - 'A']--
            l++
        }

        if r-l+1 > maxChar {
            maxChar = r-l+1
        } 
    }
    return maxChar
}
