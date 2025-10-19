func removeOccurrences(s, part string) string {
    if !strings.Contains(s, part) {
        return s // base case
    }

    // find and remove first occurrence
    for i := 0; i <= len(s)-len(part); i++ {
        if s[i:i+len(part)] == part {
            modifiedS := s[:i] + s[i+len(part):]
            return removeOccurrences(modifiedS, part)
        }
    }

    return s
}
