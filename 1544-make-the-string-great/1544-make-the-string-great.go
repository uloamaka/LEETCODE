func makeGood(s string) string {
    //with the help of rune +/- 32 of a char gives the lower or upper case
    // using a stack to keep track 
    // if stack is empty push the cur and continue
    // check the value of the char in the stack top if + 32 or -32 is equal to the next cur; pop the top and push in the next curr
    // else push the next curr
    stack := []rune{}
    for _, rn := range s {
        if len(stack) == 0 {
            stack = append(stack, rn)
            continue
        }
        top := stack[len(stack)-1]
        if top + 32  == rn || top - 32 == rn {
            stack = stack[:len(stack)-1]
            continue
        }
        stack = append(stack, rn)
    }
    return string(stack)
}