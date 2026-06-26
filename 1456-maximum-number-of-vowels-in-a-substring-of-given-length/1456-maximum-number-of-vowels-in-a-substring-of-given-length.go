func maxVowels(s string, k int) int {
    maxCnt := 0

    for i := 0; i < k; i++ {
        if isVowel(s[i]) {
            maxCnt++
        } 
    }

    if maxCnt == k {
        return maxCnt
    }

    l := 0
    vwlCnt := maxCnt
    for r := k; r < len(s); r++ {
        if maxCnt == k {
            return maxCnt
        }
        
        if isVowel(s[l]) {
            vwlCnt--
        } 
        l++
        if isVowel(s[r]) {
            vwlCnt++
        }

        if vwlCnt > maxCnt {
            maxCnt = vwlCnt
        }
    }

    return maxCnt
}

func isVowel(c byte) bool {
    switch c {
    case 'a', 'e', 'i', 'o', 'u':
        return true
    }
    return false
}
