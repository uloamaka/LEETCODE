func minAddToMakeValid(s string) int {
    // push opening parathensis to the stack and 
    // pop out from the stack only when the top of the stack is open and the current str is a closing
    // sfter that return the len of the stack
    stack := []rune{}
    // var top rune
    for _, str := range s {
        if str == '(' {
            stack = append(stack, str)
        } else if len(stack) >= 1 && stack[len(stack)-1] == '('{
            fmt.Println(stack[len(stack)-1])
            stack = stack[:len(stack)-1]
            continue
        } else {
            stack = append(stack, str)
        }   
    }
    return len(stack)
}