func findSubstring(s string, words []string) []int {
    wordLen := len(words[0])
    ans := []int{}

    required := map[string]int{}
    for _, word := range words {
        required[word]++
    }

    for offset := 0; offset < wordLen; offset++{
        window := map[string]int{}
        matched := 0

        left, right := offset, offset
        
        for right+wordLen <= len(s) {
            curWord := s[right:right+wordLen]
            right += wordLen

            if _, ok := required[curWord]; !ok {
                clear(window)
                left = right
                matched = 0
                continue
            }

            window[curWord]++
            matched++

            for window[curWord] > required[curWord] {
                leftword := s[left:left+wordLen]
                
                window[leftword]--
                matched--

                left += wordLen
            } 

            if matched == len(words) {
                ans = append(ans, left)
            }

        }

    }
    return ans
}