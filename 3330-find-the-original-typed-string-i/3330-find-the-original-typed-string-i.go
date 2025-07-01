func possibleStringCount(word string) int {
    cnt := 1
    for i := 1; i < len(word); i++ {
        if word[i-1] == word[i] {
            cnt++
        }
    }
    return cnt
}