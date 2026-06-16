func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    var s1Slc [26]int
    var s2Slc [26]int

    for i := 0; i < len(s1); i++{
        s1Slc[s1[i]-'a']++
        s2Slc[s2[i]-'a']++
    }

    if s1Slc == s2Slc {
        return true
    }

    for r := len(s1); r<len(s2); r++ {
        s2Slc[s2[r]-'a']++
        s2Slc[s2[r-len(s1)]-'a']--

        if s1Slc == s2Slc {
            return true
        }
    }

    return false
}