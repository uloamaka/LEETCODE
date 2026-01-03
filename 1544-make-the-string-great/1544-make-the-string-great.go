func makeGood(s string) string {
    stack := []rune{}
    for _, rn := range s {
        if len(stack) > 0 {
            top := stack[len(stack)-1]
            if abs(int(top - rn)) == 32{
                stack = stack[:len(stack)-1] // pop
                continue // skip current index
            }
        }
        stack = append(stack, rn) // push
    }
    return string(stack)
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}