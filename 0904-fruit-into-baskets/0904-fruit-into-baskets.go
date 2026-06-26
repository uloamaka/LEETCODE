func totalFruit(fruits []int) int {
    seen := make(map[int]int)
    var ans int

    l := 0 
    for r := 0; r < len(fruits); r++ {
        seen[fruits[r]]++

        for len(seen) > 2 {
            seen[fruits[l]]--
            if  seen[fruits[l]] == 0 {
                delete(seen, fruits[l])
            }
            l++
        }

        winSize := r-l+1
        if winSize > ans {
            ans = winSize
        }
    } 
    return ans
}