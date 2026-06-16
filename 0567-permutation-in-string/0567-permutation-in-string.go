func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    var targetSlc, windowSlc [26]int

    for i := 0; i < len(s1); i++{
        targetSlc[s1[i]-'a']++
        windowSlc[s2[i]-'a']++
    }

    if targetSlc == windowSlc {
        return true
    }

    for r := len(s1); r<len(s2); r++ {
        windowSlc[s2[r]-'a']++
        windowSlc[s2[r-len(s1)]-'a']--

        if windowSlc == targetSlc {
            return true
        }
    }

    return false
}