func maximumSum(arr []int) int {
    dp0 := arr[0]
    dp1 := 0
    res := arr[0]

    for i := 1; i < len(arr); i++ {
        // either delete current or continue deleting earlier
        dp1 = max(dp0, dp1 + arr[i]) 
        // continue or start fresh
        dp0 = max(dp0 + arr[i], arr[i]) 
          
        res = max(res, max(dp0, dp1))
    }

    return res
}