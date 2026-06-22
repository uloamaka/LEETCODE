func findAnagrams(s string, p string) []int {
    if len(p) > len(s) {
        return []int{}
    }

    var targetSlc, windowSlc [26]int
    ans := []int{}

    for i := 0; i< len(p); i++ {
        targetSlc[p[i]-'a']++
        windowSlc[s[i]-'a']++
    }

    if targetSlc == windowSlc {
        ans = append(ans, 0)
    }

    for r := len(p); r < len(s); r++ {
        windowSlc[s[r]-'a']++
        windowSlc[s[r-len(p)]-'a']--

        if targetSlc == windowSlc {
            ans = append(ans, r-len(p)+1)
        }
    }

    return ans
}