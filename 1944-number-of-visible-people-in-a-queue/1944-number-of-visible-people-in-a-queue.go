func canSeePersonsCount(heights []int) []int {
    ans := make([]int, len(heights))
    stack := []int{}
    for i := len(heights)-1; i >= 0; i-- {
        if len(stack) == 0 {
            stack = append(stack, heights[i])
            continue
        }
        // it should be recursive here
        for len(stack) > 1 && heights[i] > stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
            ans[i] += 1
        }

        if heights[i] > stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
        }
        
        stack = append(stack, heights[i])
        ans[i] += 1
    }
    return ans
}