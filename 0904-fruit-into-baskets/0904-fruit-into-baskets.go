func totalFruit(fruits []int) int {
    fruitsM := make(map[int]int)
    start := 0
    maxFruits := 0

    for end := 0; end < len(fruits); end++ {
        fruitsM[fruits[end]]++

        if len(fruitsM) > 2 {
            fruitsM[fruits[start]]--

            if fruitsM[fruits[start]] == 0 {
                delete(fruitsM, fruits[start])
            }
            start++
        }
        
        windowSize := end - start + 1
        if windowSize > maxFruits {
            maxFruits = windowSize
        }
    }

    return maxFruits
}