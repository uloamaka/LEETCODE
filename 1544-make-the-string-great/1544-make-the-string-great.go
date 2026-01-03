func makeGood(s string) string {
    stack := []rune{}
    for _, rn := range s {
        if len(stack) > 0 {
            top := stack[len(stack)-1]
            if top + 32  == rn || top - 32 == rn {
                stack = stack[:len(stack)-1] // pop
                continue
            }
        }
        stack = append(stack, rn)
    }
    return string(stack)
}