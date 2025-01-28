func getSneakyNumbers(nums []int) []int {
    counts := make(map[int]int)
    sneakyNums := []int{}
    for _, num := range nums {
        counts[num]++
    }
    for num, count := range counts {
        if count > 1 {
            sneakyNums = append(sneakyNums, num)
        }
    }
    return sneakyNums
}

// count each unique numbers in a given array as N,
// return unique numbers with counts more than 1.