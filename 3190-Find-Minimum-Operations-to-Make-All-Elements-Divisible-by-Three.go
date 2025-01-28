func minimumOperations(nums []int) int {
    cnt := 0
    for _, num := range nums {
        if num % 3 != 0 {
            cnt++ 
        }
    }
    return cnt
}
// loop through the slice 
// If x % 3 != 0 then count x