func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    keyIndex := map[string]int{"type": 0, "color": 1, "name": 2}
    index := keyIndex[ruleKey] 

    cnt := 0
    for _, item := range items {
        if item[index] == ruleValue {
            cnt++
        }
    }
    return cnt
}