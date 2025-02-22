func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandy := 0
    for _, candy := range candies {
        if candy > maxCandy {
            maxCandy = candy
        }
    }
    
    result := make([]bool, len(candies))
    for i, candy := range candies {
        result[i] = (candy + extraCandies) >= maxCandy
    }
    
    return result
}