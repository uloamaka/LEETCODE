func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    var str1, str2 strings.Builder

    for _, val := range word1 {
        str1.WriteString(val)
    }
    for _, val := range word2 {
        str2.WriteString(val)
    }

    return str1.String() == str2.String()
}