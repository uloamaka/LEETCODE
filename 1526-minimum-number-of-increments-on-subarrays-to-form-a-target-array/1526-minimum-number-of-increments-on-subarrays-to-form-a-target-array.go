func minNumberOperations(target []int) int {
    stack := []int{}
    ops := 0

    for _, h := range target {
        if len(stack) == 0 {
            ops += h
        } else if stack[len(stack)-1] < h {
            ops += h - stack[len(stack)-1]
        }

        stack = append(stack, h)
    }

    return ops
}
