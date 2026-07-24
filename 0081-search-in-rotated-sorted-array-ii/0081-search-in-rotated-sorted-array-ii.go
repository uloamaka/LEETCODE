func search(nums []int, target int) bool {
    // left, right := 0, len(nums)-1

    // for left < right {
    //     mid := left + (right-left) / 2

    //     if nums[mid] == target {
    //         return true
    //     }

    //     // [left, mid] is sorted
    //     if nums[left] <= nums[mid] {
    //         if nums[left] <= target && target < nums[mid] {
    //             right = mid - 1
    //         } else {
    //             left = mid + 1
    //         }
    //     } else {
    //         if nums[mid] < target && target <= nums[right] {
    //             left = mid + 1
    //         } else {
    //             right = mid - 1
    //         }
    //     }
    // }

    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            return true
        }
    }
    return false
}
