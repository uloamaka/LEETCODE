func findMaxAverage(nums []int, k int) float64 {
    var sum int
 
    for i := 0; i < k; i++ {
        sum += nums[i]
    }

    maxSum := sum

    for right := k; right < len(nums); right++ {
        sum += nums[right] - nums[right - k]

        if sum > maxSum {
            maxSum = sum
        }
    }

    return float64(maxSum) / float64(k)
}

func init(){
    debug.SetMemoryLimit(0)
}
