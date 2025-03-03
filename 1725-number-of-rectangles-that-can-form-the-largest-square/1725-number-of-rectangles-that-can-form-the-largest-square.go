func countGoodRectangles(rectangles [][]int) int {
    m := make(map[int]int)
    maxLen := 0

    for _, rectangle := range rectangles {
        minLen := min(rectangle[0], rectangle[1])
        m[minLen]++
        
        if minLen > maxLen {
            maxLen = minLen
        }
    }

    return m[maxLen]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

