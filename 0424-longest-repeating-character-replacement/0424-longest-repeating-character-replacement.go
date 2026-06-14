func characterReplacement(s string, k int) int {
    freq := map[byte]int{}
    maxChar, curLen, l := 0, 0, 0

    for r := 0; r < len(s); r++ {
        freq[s[r]]++

        maxfreq := maxFreq(freq)
        curLen = r - l + 1
        
        if curLen - maxfreq  > k {
            freq[s[l]]--
            l++
            curLen = r-l+1
        }

        if curLen > maxChar {
            maxChar = curLen
        } 
    }
    return maxChar
}

func maxFreq(m map[byte]int) int {
    max := 0
    for _, cnt := range m {
        if cnt > max {
            max = cnt
        }
    }
    return max
}