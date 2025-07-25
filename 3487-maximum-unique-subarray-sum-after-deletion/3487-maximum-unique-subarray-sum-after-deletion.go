func maxSum(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    var ans int
    var maxNegative int
    maxNum := -100
    uniqueM := make(map[int]bool)

    for _, num := range nums {
        if num > maxNum {
            maxNum = num
        }
        maxNegative = maxNum

        _, ok := uniqueM[num]
        if !ok && num >= 0 {
            uniqueM[num] = true 
            ans += num
        }
    }
    if maxNegative < 0 {
        return maxNegative
    }

    return ans
}
    // Give priority to deleting -ve numbers, (if a num is -ve, )
    // use a hashmap to check for unique elements, if a num is unique add it to ans.