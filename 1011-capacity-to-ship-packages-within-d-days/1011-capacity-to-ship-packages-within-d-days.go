func shipWithinDays(weights []int, days int) int {
    maxWeight, totalWeight := 0, 0
    for _, weight := range weights {
        totalWeight += weight

        if weight > maxWeight {
            maxWeight = weight
        }
    }

    minCap, maxCap := maxWeight, totalWeight
    for minCap < maxCap {
        mid := minCap + (maxCap-minCap) / 2

        estDays := getDays(weights, mid)
        if estDays <= days {
            maxCap = mid
        } else {
            minCap = mid+1
        }
    }
    return minCap
}

func getDays(weights []int, shpCap int) int {
    summer := 0
    cnt := 1
    for _, weight := range weights {
        if summer + weight > shpCap {
            cnt++
            summer = weight
        } else {
            summer += weight
        }
    }
    return cnt
}
