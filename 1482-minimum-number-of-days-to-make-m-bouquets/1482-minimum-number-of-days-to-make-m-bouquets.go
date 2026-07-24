func minDays(bloomDay []int, m int, k int) int {
    if m*k > len(bloomDay) {
        return -1
    }

    left, right := 0, slices.Max(bloomDay)
    minDays := -1

    for left <= right {
        mid := left + (right-left) / 2

        if getNumOfBouquets(bloomDay, mid, k) >= m {
            minDays = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return minDays
}

func getNumOfBouquets(bloomDay []int, mid int, k int) int {
    numOfBouquets := 0
    cnt := 0
    for i := 0; i < len(bloomDay); i++ {
        if bloomDay[i] <= mid {
            cnt++
        } else {
            cnt = 0
        }
        if cnt == k {
            numOfBouquets++
            cnt = 0
        }
    } 
    return numOfBouquets
}