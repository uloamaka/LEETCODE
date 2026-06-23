func maxArea(height []int) int {
    l, r := 0, len(height)-1
    maxArea := 0

    for l < r {
        width, area := 0, 0
        if height[l] < height[r] {
            width = r-l
            area = width * height[l]
            l++
        } else if height[l] == height[r] {
            width = r-l
            area = width * height[l]
            r--
        } else  {
            width = r-l
            area = width * height[r]
            r--
        }
        maxArea = max(maxArea, area)
    }
    return maxArea
}
