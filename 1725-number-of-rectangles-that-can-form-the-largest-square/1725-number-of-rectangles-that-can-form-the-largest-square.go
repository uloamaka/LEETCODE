func countGoodRectangles(rectangles [][]int) int {
    maxLens := make([]int, len(rectangles))
    m := make(map[int]int)
    
    for i, rectangle := range rectangles {
        minLen := slices.Min(rectangle)
        maxLens[i] = minLen

        m[minLen] += 1
    }

    maxLen := slices.Max(maxLens)
    return m[maxLen]
}
