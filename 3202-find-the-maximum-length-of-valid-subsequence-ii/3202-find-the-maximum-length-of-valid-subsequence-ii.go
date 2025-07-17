func maximumLength(nums []int, k int) int {
    res := 1 

    for m := 0; m < k; m++ {
        dp := make(map[int]int)
        for _, num := range nums {
            cur := num % k
            prev := (m - cur + k) % k
            dp[cur] = dp[prev] + 1
            res = max(res, dp[cur])
        }
    }

    return res
}