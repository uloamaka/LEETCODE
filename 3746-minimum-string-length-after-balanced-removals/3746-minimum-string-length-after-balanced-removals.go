func minLengthAfterRemovals(s string) int {
    stack := []rune{}

    for _, rn := range s {
        if len(stack) > 0 && rn == 'b' && stack[len(stack)-1] == 'a' {
            stack = stack[:len(stack)-1]
            continue
        }

        if len(stack) > 0 && rn == 'a' && stack[len(stack)-1] == 'b' {
            stack = stack[:len(stack)-1]
            continue
        }
        stack = append(stack, rn)
    }
    return len(stack)
}
