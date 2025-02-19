func minOperations(nums []int, k int) int {
    ans := 0
    sort.Ints(nums)
    for i, val := range nums {
        if val >= k {
            ans = i
            break
        }
    }
    return ans
}
// Algorithm
// sort the nums array
// loop through the nums 
// range through the slice get the first index where the value is >= to k
// then return that index