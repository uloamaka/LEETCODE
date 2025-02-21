func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    i, j := 0, 0 
    ci, cj := 0, 0 

    for i < len(word1) && j < len(word2) {
        if word1[i][ci] != word2[j][cj] {
            return false
        }

        ci++
        cj++

        if ci == len(word1[i]) {
            i++
            ci = 0
        }

        if cj == len(word2[j]) {
            j++
            cj = 0
        }
    }

    return i == len(word1) && j == len(word2)
}
