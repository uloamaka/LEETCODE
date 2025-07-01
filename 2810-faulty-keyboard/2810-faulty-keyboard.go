func finalString(s string) string {
    var str strings.Builder
    i := 0
    
    for i < len(s) {
        if s[i] == 'i' {
            // Reverse everything we've built so far
            current := str.String()
            str.Reset()
            str.WriteString(reverse(current))
            i++ // Skip the 'i'
        } else {
            str.WriteByte(s[i]) // Add current character to our string
            i++
        }
    }
    
    return str.String()
}

func reverse(s string) string {
    r := []rune(s) // Convert string to a slice of runes
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i] // Swap runes
    }
    return string(r) // Convert rune slice back to string
}

// Intuition
    // loop thro the s string
    // anytime "i" is found call a reverse function on the index of the runes b4 'i'
    // return the final string