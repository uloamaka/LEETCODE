func countGoodSubstrings(s string) int {
    if len(s) < 3 {
        return 0
    }

    seen := make(map[byte]int)
    distinct := 0

    for i := 0; i < 3; i++ {
        if seen[s[i]] == 0 {
            distinct++
        }
        seen[s[i]]++
    }

    ans := 0
    if distinct == 3 {
        ans++
    }

    for r := 3; r < len(s); r++ {
        l := r - 3

        seen[s[l]]--
        if seen[s[l]] == 0 {
            distinct--
        }

        if seen[s[r]] == 0 {
            distinct++
        }
        seen[s[r]]++

        if distinct == 3 {
            ans++
        }
    }

    return ans
}