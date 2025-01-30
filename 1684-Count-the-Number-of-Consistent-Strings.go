func countConsistentStrings(allowed string, words []string) int {
    allowedStrCnt := 0
    for _, word := range words {
        for i, v := range word {
            if ok := strings.Contains(allowed, string(v)); !ok {
                break
            }
            if i == len(word) - 1 {
                 allowedStrCnt++
            }
        }
    }
    return allowedStrCnt
}