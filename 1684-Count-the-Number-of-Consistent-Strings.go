func countConsistentStrings(allowed string, words []string) int {
    allowedStrCnt := 0
    for _, word := range words {
        isValid := true
        for _, v := range word {
            if !strings.Contains(allowed, string(v)) {
                isValid = false
                break
            }
        }
        if isValid {
            allowedStrCnt++
        }
    }
    return allowedStrCnt
}