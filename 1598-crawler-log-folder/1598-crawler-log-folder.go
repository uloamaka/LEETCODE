func minOperations(logs []string) int {
    stack := []int{}

    for _, vl := range logs {
        if vl == "../" {
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
            continue
        }

        if vl == "./" {
            continue
        }
        
        stack = append(stack, 1)
    }

    return len(stack)
}