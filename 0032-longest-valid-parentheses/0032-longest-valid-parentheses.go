func longestValidParentheses(s string) int {
    stack := []int{}
    base := -1
    maxLen := 0

    for i, rn := range s {
        if rn == '(' {
            stack = append(stack, i)
        } else {
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]

                var length int

                if len(stack) > 0 {
                    length = i - stack[len(stack)-1]
                } else {
                    length = i - base
                } 
            
                if length > maxLen {
                    maxLen = length
                }
            } else {
                base = i
            }
        }
    }
    return maxLen
}