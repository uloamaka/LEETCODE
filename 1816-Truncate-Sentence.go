func truncateSentence(s string, k int) string {
    slcStr := strings.Split(s, " ")

    if k < len(slcStr) {
        return strings.Join(slcStr[:k], " ")
    } else {
        return s
    }
}