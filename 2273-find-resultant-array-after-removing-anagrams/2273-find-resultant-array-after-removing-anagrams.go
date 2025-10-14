func removeAnagrams(words []string) []string {
    stack := []string{words[0]}

    top := words[0]

    for i := 0; i < len(words); i++  {
        top = stack[len(stack)-1]
        if !isAnagram(top, words[i]) {
            stack = append(stack, words[i])
        }
    }

    return stack
}

func isAnagram(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }

    freq := [26]int{}
    for i := range s1 {
        freq[s1[i]-'a']++
        freq[s2[i]-'a']--
    }

    for _, v := range freq {
        if v != 0 {
            return false
        }
    }

    return true
}

// func isAnagram(s1, s2 string) bool {
//     freqM1 := make(map[rune]int)
//     freqM2 := make(map[rune]int)

//     for _, rn := range s1 {
//         freqM1[rn]++
//     }

//     for _, rn := range s2 {
//         freqM2[rn]++
//     }

//     if len(freqM1) != len(freqM2) {
//         return false
//     }

//     for r, cnt := range freqM1 {
//         if freqM2[r] != cnt {
//             return false
//         }
//     }

//     return true
// }

// write a func is anagram,
// the first word is compared with the  next word, 