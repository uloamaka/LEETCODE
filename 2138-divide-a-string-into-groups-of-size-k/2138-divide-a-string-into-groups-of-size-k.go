func divideString(s string, k int, fill byte) []string {
    totalLen := len(s)
    result := []string{}
    fullGroups := totalLen / k
    startIdx, endIdx := 0, k

    // Add full groups
    for i := 0; i < fullGroups; i++ {
        result = append(result, s[startIdx:endIdx])
        startIdx += k
        endIdx += k
    }

    // Handle last group if it's not a full group
    remaining := totalLen % k
    if remaining != 0 {
        lastGroup := s[totalLen-remaining:]
        for i := 0; i < k-remaining; i++ {
            lastGroup += string(fill)
        }
        result = append(result, lastGroup)
    }

    return result
}