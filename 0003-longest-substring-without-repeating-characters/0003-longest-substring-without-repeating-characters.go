func lengthOfLongestSubstring(s string) int {
    dict := make([]int, 128)
    var  maxLen, l int
    for r := 0; r < len(s); r++{
        currChar := s[r]
        if dict[currChar] > l {
            l = dict[currChar]
        }
        if r - l + 1 > maxLen{
            maxLen = r - l + 1
        }
        dict[currChar] = r + 1
    }
    return maxLen
}
