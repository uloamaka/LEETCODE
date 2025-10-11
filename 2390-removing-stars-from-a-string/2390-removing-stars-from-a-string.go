func removeStars(s string) string {
    stack := []rune{}
    for _, rn := range s {
        if rn != '*' {
            stack = append(stack, rn)
        } else if len(stack) != 0 {
            stack = stack[:len(stack)-1]
        }
    }
    return string(stack)
}