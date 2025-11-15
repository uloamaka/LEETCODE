func minLength(s string) int {
    stack := []byte{}

    for i := 0; i < len(s); i++ {
        if len(stack) == 0 {
            stack = append(stack, s[i])
            continue
        }
        if s[i] == 'B' && stack[len(stack)-1] == 'A' {
            stack = stack[:len(stack)-1]
        } else if s[i] == 'D' && stack[len(stack)-1] == 'C' {
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, s[i])
        }
    }
    return len(stack)
}