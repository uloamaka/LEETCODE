func minWindow(s string, t string) string {
    if len(t) > len(s) {
        return ""
    }

    required := map[byte]int{}
    for i := 0; i<len(t); i++ {
        required[t[i]]++
    }

    start := 0
    minLen := math.MaxInt32
    formed := 0
    requiredUnique := len(required)

    l := 0
    window := map[byte]int{}
    for r := 0; r < len(s); r++ {
        window[s[r]]++

        if _, ok := required[s[r]]; ok && required[s[r]] == window[s[r]] {
            formed++
        }

        for formed == requiredUnique {
            if r-l+1 < minLen {
                minLen = r-l+1
                start = l
            }

            window[s[l]]--
            if window[s[l]] < required[s[l]] {
                formed--
            }
            l++
        }
    }

    if minLen == math.MaxInt32 {
        return ""
    }

    return s[start : start+minLen]
}