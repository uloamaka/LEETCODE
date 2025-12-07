func minAddToMakeValid(s string) int {
    stack := []rune{}
    for _, str := range s {
        if str == '(' {
            stack = append(stack, str)
        } else if len(stack) >= 1 && stack[len(stack)-1] == '('{
            stack = stack[:len(stack)-1]
            continue
        } else {
            stack = append(stack, str)
        }   
    }
    return len(stack)
}