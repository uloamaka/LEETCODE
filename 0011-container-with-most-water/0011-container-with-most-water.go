func maxArea(height []int) int {
    maxArea := 0
    l, r := 0, len(height)-1

    for l < r {
        h := 0
        if height[l] < height[r] {
            h = height[l]
            l++
        } else {
            h = height[r]
            r--
        }

        area := h * (r - l + 1)  // +1 because l or r has already been moved
        if area > maxArea {
            maxArea = area
        }
    }

    return maxArea
}
