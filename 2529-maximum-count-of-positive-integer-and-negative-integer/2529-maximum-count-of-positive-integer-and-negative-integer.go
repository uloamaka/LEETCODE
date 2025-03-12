func maximumCount(nums []int) int {
    n := len(nums)
    pCnt := n - geNonNegative(nums, n)
    nCnt := getFirstPositive(nums, n)

    if nCnt > pCnt {
        return nCnt
    }
    
    return pCnt 
}

func geNonNegative (nums []int, n int) int {
    r, l:= n, 0
    for l < r {
        mid := l + (r-l)/2
        if nums[mid] > 0 {
            r = mid 
        } else {
            l = mid + 1
        }
    }
    return l
}
func getFirstPositive (nums []int, n int) int {
    r, l:= n, 0
    for l < r {
        mid := l + (r-l)/2
        if nums[mid] >= 0 {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return l
}