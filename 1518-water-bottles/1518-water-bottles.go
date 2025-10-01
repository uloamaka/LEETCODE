func numWaterBottles(numBottles int, numExchange int) int {
    totalDrank, empty := numBottles, numBottles
    for empty >= numExchange {
        totalDrank += 1
        empty -= numExchange-1
    }
    return totalDrank
}