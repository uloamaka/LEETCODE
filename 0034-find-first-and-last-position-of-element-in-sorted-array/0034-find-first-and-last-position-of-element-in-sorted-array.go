func searchRange(nums []int, target int) []int {
    ans := []int{-1,-1}
    left, right := 0, len(nums)
    for left < right {
        mid := left + (right-left) / 2
        
        if nums[mid] == target {
            ans[0] = mid
        }

        if nums[mid] < target {
            left = mid+1
        } else {
            right = mid
        }
    }

    left2, right2 := 0, len(nums)
    for left2 < right2 {
        mid := left2 + (right2-left2) / 2

        if nums[mid] == target {
            ans[1] = mid
        } 
        if nums[mid] <= target {
            left2 = mid+1
        } else {
            right2 = mid
        }
    }
    return ans
}