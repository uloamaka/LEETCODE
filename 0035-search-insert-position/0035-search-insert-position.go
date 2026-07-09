func searchInsert(nums []int, target int) int {
    left, right, lastChk := 0, len(nums), 0

    for left < right {
        mid := left + (right - left) / 2

        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
            lastChk = mid+1
        } else {
            right = mid
            lastChk = mid
        }
    }
    return lastChk
}