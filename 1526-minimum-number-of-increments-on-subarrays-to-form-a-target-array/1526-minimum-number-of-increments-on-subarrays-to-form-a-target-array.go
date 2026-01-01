func minNumberOperations(target []int) int {
    cnt := target[0]
    lastSeen := target[0]
    for i := 1; i < len(target); i++ {
        if target[i] <= lastSeen {
            lastSeen = target[i]
            continue
        }
        toAdd := target[i] - lastSeen
        cnt += toAdd
        lastSeen = target[i]
    }
    return cnt
}