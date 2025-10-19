func removeOccurrences(s, part string) string {
    stack := []byte{}
    n := len(part)

    for i := 0; i < len(s); i++ {
        stack = append(stack, s[i])
        if len(stack) >= n && string(stack[len(stack)-n:]) == part {
            stack = stack[:len(stack)-n]
        }
    }

    return string(stack)
}
