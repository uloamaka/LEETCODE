// func truncateSentence(s string, k int) string {
//     slcStr := strings.Split(s, " ")

//     if k < len(slcStr) {
//         return strings.Join(slcStr[:k], " ")
//     } else {
//         return s
//     }
// }

func truncateSentence(s string, k int) string {
    space := rune(' ')
    spaceCnt := 0

    for idx, char := range s {
        if char == space {
            spaceCnt++
        }
        if k == spaceCnt {
            return s[0 : idx]
        }
    } 
    return s
}