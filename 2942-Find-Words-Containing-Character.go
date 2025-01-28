func findWordsContaining(words []string, x byte) []int {
    wordsIndex := []int{}
    xStr := string(x)
    for i, word := range words {
        if strings.Contains(word, xStr) {
            wordsIndex = append(wordsIndex, i)
        }
    }
    return wordsIndex
}