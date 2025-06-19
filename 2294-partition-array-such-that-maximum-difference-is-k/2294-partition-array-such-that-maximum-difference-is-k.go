func partitionArray(nums []int, k int) int {
    sort.Ints(nums)
    cnt := 0
    left := 0
    
    for right := 0; right < len(nums); right++ {
        if nums[right]-nums[left] > k {
            cnt++
            left = right
        }
    }
    return cnt + 1
}