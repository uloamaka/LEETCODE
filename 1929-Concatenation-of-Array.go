func getConcatenation(nums []int) []int {
// Determine the length of the input array
// Create a new array with twice the original length
// Copy the original array's elements to the first half of the new array
// Copy the original array's elements again to the second half of the new array
    arrLen := len(nums)
    ans := make([]int, arrLen * 2)

    for i, num := range nums {
        ans[i] = num
        ans[i + arrLen] = num
    }

    return ans
}