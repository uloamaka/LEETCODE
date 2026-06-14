func characterReplacement(s string, k int) int {
    var freq [256]int
    var maxChar int
    var maxFreq int 
    
    var l int
    for r := 0; r < len(s); r++ {

        freq[s[r]]++

        maxFreq = max(maxFreq, freq[s[r]])

        if r - l + 1 - maxFreq  > k {
            freq[s[l]]--
            l++
        }

        if r-l+1 > maxChar {
            maxChar = r-l+1
        } 
    }
    return maxChar
}
