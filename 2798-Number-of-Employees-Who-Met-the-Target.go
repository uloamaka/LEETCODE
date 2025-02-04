func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    ans := 0
    for _, hour := range hours {
        if hour >= target {
            ans++
        }
    }
    return ans
}
