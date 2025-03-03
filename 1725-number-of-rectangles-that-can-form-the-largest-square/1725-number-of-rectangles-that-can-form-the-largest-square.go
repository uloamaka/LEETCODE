func countGoodRectangles(rectangles [][]int) int {
    maxLens := make([]int, len(rectangles))
    for i, rectangle := range rectangles {
        min := slices.Min(rectangle)
        maxLens[i] = min
    }

    maxINT := slices.Max(maxLens)
    cnt := 0
    for _, len := range maxLens {
        if len == maxINT {
            cnt++
        }
    }
    return cnt
}
