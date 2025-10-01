func numWaterBottles(numBottles int, numExchange int) int {
    ttl, emp := numBottles, numBottles
    for emp >= numExchange {
        ttl += 1
        emp -= numExchange-1
    }
    return ttl
}