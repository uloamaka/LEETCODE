func trap(height []int) int {
    stack := []int{}
    water := 0 

    for i := range height {
        for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
            bottom := stack[len(stack)-1] // lowest height 
            bottomHeight := height[bottom]
            stack = stack[:len(stack)-1] // pop off from height index stack 

            if len(stack) == 0 { // empty stack safety
                break
            }

            leftWall := height[stack[len(stack)-1]] 
            rightWall := height[i] 

            width := i - stack[len(stack)-1] - 1
            height := min(rightWall, leftWall) - bottomHeight
            water += height * width
        } 

        stack = append(stack, i)
    }
    
    return water
}

func init() {
    debug.SetMemoryLimit(10)
}