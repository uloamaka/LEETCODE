func findMaxAverage(nums []int, k int) float64 {
    var maxSum, sum float64
 
    for i := 0; i < k; i++ {
        sum += float64(nums[i])
    }
    maxSum = sum

    for right := k; right < len(nums); right++ {
        sum = sum - float64(nums[right - k])

        sum += float64(nums[right])

        if sum > maxSum {
            maxSum = sum
        }
    }
    
    maxAvg := maxSum/float64(k)
    return maxAvg
}
