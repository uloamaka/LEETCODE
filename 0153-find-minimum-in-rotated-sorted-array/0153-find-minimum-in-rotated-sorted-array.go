func findMin(nums []int) int {
    n := len(nums)-1
    left, right := 0, n

    for left < right {
        mid := left + (right- left) / 2

        if nums[mid] <= nums[n] {
            right = mid
        } else {
            left = mid+1
        }
    }
    return nums[left]
}