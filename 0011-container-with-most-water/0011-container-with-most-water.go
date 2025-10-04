func maxArea(height []int) int {
    // fmt.Println(currHeight)
    maxArea := 0 
    l, r := 0, len(height)-1
    for l <= r {
        currHeight, currBase := 0, 0
        if height[l] < height[r] {
            currHeight = height[l]
            currBase = AbsInt(l - r)
            // fmt.Println(currHeight*currBase)
            // // fmt.Println(currBase)
            l++
        } else if height[l] == height[r] {
            currHeight = height[r]
            currBase = AbsInt(l - r)
            // fmt.Println(currHeight*currBase)
            // fmt.Println(currBase)
            r--
        } else {
            currHeight = height[r]
            currBase = AbsInt(l - r)
            // fmt.Println(currHeight*currBase)
            // fmt.Println(currBase)
            r--
        }
    
        area := currHeight*currBase
        if area > maxArea {
            maxArea = area
        }

    }
    return maxArea
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}