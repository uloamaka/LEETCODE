func countMaxOrSubsets(nums []int) int {
    maxOR := 0
    for _, x := range nums {
        maxOR |= x
    }

    count := 0

    var backtrack func(int, int)
    backtrack = func(i, currOR int) {
        if i == len(nums) {
            if currOR == maxOR {
                count++
            }
            return
        }

        backtrack(i+1, currOR|nums[i])
        backtrack(i+1, currOR)
    }

    backtrack(0, 0)
    return count
}