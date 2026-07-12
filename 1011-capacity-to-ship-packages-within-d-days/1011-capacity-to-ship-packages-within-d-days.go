func shipWithinDays(weights []int, days int) int {
    maxWeight, totalWeight := 0, 0
    for _, weight := range weights {
        totalWeight += weight

        if weight > maxWeight {
            maxWeight = weight
        }
    }

    left, right := maxWeight, totalWeight
    for left < right {
        mid := left + (right-left) / 2

        estDays := daysNeeded(weights, mid)
        if estDays <= days {
            right = mid
        } else {
            left = mid+1
        }
    }
    return left
}

func daysNeeded(weights []int, capacity int) int {
    currentLoad := 0
    days := 1
    for _, weight := range weights {
        if currentLoad + weight > capacity {
            days++
            currentLoad = weight
        } else {
            currentLoad += weight
        }
    }
    return days
}

// Time: O(n log W)
// Space: O(1)
