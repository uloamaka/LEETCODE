func makeFancyString(s string) string {
    var sb strings.Builder
    count := 1 // Track consecutive characters

    sb.WriteByte(s[0])

    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {
            count++
        } else {
            count = 1
        }

        if count <= 2 {
            sb.WriteByte(s[i])
        }
    }

    return sb.String()
}