func numIdenticalPairs(nums []int) int {
    counts := make(map[int]int)
    sum := 0
    for _, num := range nums {
        counts[num]++
    }
    for _, count := range counts {
       sum += count * (count - 1)/2
    }
    return sum
}
    // well I didn't understand the question
    // but here is what i got for the discussion segment
    // count each unique numbers in a given array as N, 
    // then applly the formula N * (N - 1) / 2 
    // then sum the total count of all N