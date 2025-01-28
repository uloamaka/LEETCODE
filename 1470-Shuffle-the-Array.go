func shuffle(nums []int, n int) []int {
    slice1 := nums[:n]
    slice2 := nums[n:]
    shuffledSlc := []int{}
    for i := 0; i < n; i++ {
        shuffledSlc = append(shuffledSlc, slice1[i], slice2[i] )
    }
    return shuffledSlc
}