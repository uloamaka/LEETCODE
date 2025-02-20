func mostWordsFound(sentences []string) int {
    wordCnt := 0
    for _, sentence := range sentences {
        nWord := 0
        for _, char := range sentence {
            if char == ' ' {
               nWord += 1
            }
        }
        nWord += 1 
        if wordCnt < nWord {
            wordCnt = nWord
        }
    }
    
    return wordCnt
}
