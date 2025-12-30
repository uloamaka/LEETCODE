func validateStackSequences(pushed []int, popped []int) bool {
    n := len(pushed)
    idxPop := 0
    idxPush := 0
    stack := []int{}

    for idxPush < n {
        stack = append(stack, pushed[idxPush])
        idxPush++

        for idxPop < n && len(stack) > 0 && popped[idxPop] == stack[len(stack)-1] {
            idxPop++
            stack = stack[:len(stack)-1]
        }     
    }
    return len(stack) == 0 
}
