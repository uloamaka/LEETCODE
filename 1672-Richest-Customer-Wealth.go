func maximumWealth(accounts [][]int) int {
    maxBal := 0
    
    for _, acc := range accounts {
        currentBal := 0
        for _, amount := range acc {
            currentBal += amount
        }
        if currentBal > maxBal {
            maxBal = currentBal
        }
    }
    
    return maxBal
}
