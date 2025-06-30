func findLHS(nums []int) int {
    freq := make(map[int]int)
    for _, num := range nums {
        freq[num]++
    }

    keySli := []int{}
    for key, _ := range freq {
        keySli = append(keySli, key)
    }

    highestFreqCnt := 0
    for i := 1; i < len(keySli); i++ {
        if keySli[i]-keySli[i-1] == 1 {
            sum := freq[keySli[i]] + freq[keySli[i-1]]
            if sum > highestFreqCnt {
                highestFreqCnt = sum
            }
        }
    }

    return highestFreqCnt
}