func search(nums []int, target int) int {
    left := 0
    right := len(nums) 

    for left < right {
        mid := left + (right - left) / 2
        if nums[mid] <= target {
            left = mid + 1
        } else {
            right = mid
        } 
    } 
    if left > 0 && nums[left - 1] == target{
        return  left - 1
    }  
    return -1 
}

    // for left <= right {
    //     mid := left + (right - left) / 2
    //     if nums[mid] == target {
    //         return mid
    //     } else if nums[mid] < target {
    //         left = mid + 1
    //     } else {
    //         right = mid - 1
    //     }
    // }
