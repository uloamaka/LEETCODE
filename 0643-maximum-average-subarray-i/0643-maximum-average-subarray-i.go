func findMaxAverage(nums []int, k int) float64 {
    var maxAvg, sum float64
    var left int

    for i := 0; i < k; i++ {
        sum += float64(nums[i])
        maxAvg = sum /float64(k)
    }

    for right := k; right < len(nums); right++ {
        sum = sum - float64(nums[left])

        sum += float64(nums[right])
        left++

        avg := sum/float64(k)
        if avg > maxAvg {
            maxAvg = avg
        }
    }
    return maxAvg
}
