func findPeakElement(nums []int) int {
    n := len(nums)
    if n == 1 {
        return 0
    }

    if nums[0] > nums[1] {
        return 0
    }

    if nums[n-1] > nums[n-2] {
        return n - 1
    }

    left, right := 1, len(nums)-2
    for left <= right {
        mid := left + (right - left) / 2

        if nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
            return mid
        } 

        if nums[mid] < nums[mid+1] {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
   
    return -1
}