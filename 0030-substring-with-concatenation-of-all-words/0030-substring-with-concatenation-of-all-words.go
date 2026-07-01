func findSubstring(s string, words []string) []int {
    wordLen := len(words[0])
    ans := []int{}

    required := make(map[string]int)
    for _, word := range words {
        required[word]++
    }

    for offset := 0; offset < wordLen; offset++ {
        left, right := offset, offset
        window := map[string]int{}
        matched := 0

        for right+wordLen <= len(s) {
            word := s[right:right+wordLen]
            right += wordLen

            window[word]++
            matched++

            if _, ok := required[word]; !ok {
                clear(window)
                matched = 0
                left = right
                continue
            }

            for window[word] > required[word] {
                leftWord := s[left:left+wordLen]

                window[leftWord]--
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