func decompressRLElist(nums []int) []int {
    size := 0
    for i := 0; i < len(nums); i += 2 {
        size += nums[i] 
    }

    ans := make([]int, size)
    index := 0

    for i := 0; i < len(nums); i += 2 {
        freq, val := nums[i], nums[i+1]
        for j := 0; j < freq; j++ {
            ans[index] = val
            index++
        }
    }
    return ans
}