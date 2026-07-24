func binarySearch(nums []int, left, right, target int) int{
    if left > right {
        return -1
    }

    mid := left + (right-left)/2
    if nums[mid] == target {
        return mid
    }

    // left - mid part is sorted
    if nums[left] <= nums[mid] {
        if target >= nums[left] && target < nums[mid] {
            return binarySearch(nums, left, mid-1, target)
        }
        return binarySearch(nums, mid+1, right, target)
    } else {
        // mid - right is sorted
        if target <= nums[right] && nums[mid] < target {
            return binarySearch(nums, mid+1, right, target)   
        }
        return binarySearch(nums, left, mid-1, target)
    }
}


func search(nums []int, target int) int {
    return binarySearch(nums, 0, len(nums)-1, target)
}