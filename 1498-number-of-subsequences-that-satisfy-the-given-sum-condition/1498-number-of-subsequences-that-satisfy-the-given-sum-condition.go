func numSubseq(nums []int, target int) int {
    const mod = 1e9 + 7
    n := len(nums)

    sort.Ints(nums)

    power2 := make([]int, n+1)
    power2[0] = 1
    for i := 1; i <= n; i++ {
        power2[i] = (power2[i-1] * 2) % mod
    }

    ans := 0
    left, right := 0, n-1
    for left <= right {
        if nums[left]+nums[right] <= target {
            ans = (ans + power2[right-left]) % mod
            left++
        }else {
            right--
        }
    }
    return ans
}