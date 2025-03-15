func numberOfSubstrings(s string) int {
    len := len(s)
    left, right, total := 0, 0, 0
    freq := make([]int, 3)
    for right < len {
       freq[s[right] - 'a']++
        fmt.Println(freq)
       for freq[0] > 0 && freq[1] > 0 && freq[2] > 0 {
            total += len - right
            freq[s[left] - 'a']--
            left++
       }

       right++
    }
    return total   
}
