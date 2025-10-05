func maxDepth(s string) int {
    stack := []rune{}

    maxSeen := 0
    for _, rn := range s {

        if rn == '(' {
            stack  = append(stack, rn)
        }
        
        if rn == ')' {
            // if stack[len(stack)-1] == '(' {
                if len(stack) > maxSeen {
                    maxSeen = len(stack)
                }
                stack = stack[:len(stack)-1]
            // }
        }
    }
    return maxSeen
}