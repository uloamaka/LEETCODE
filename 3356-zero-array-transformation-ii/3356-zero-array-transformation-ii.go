func minZeroArray(nums []int, queries [][]int) int {
    left, right := 0, len(queries)
    result := -1

    for left <= right {
        mid := left + (right - left) / 2
        if canFormZeroArray(nums, queries, mid) {
            result = mid
            right = mid - 1 // Look for a smaller k
        } else {
            left = mid + 1
        }
    }

    return result
}

func canFormZeroArray(nums []int, queries [][]int, k int) bool {
    n := len(nums)
    differenceArray := make([]int, n + 2) // n+2 to avoid index issues

    // Apply first k queries
    for i := 0; i < k; i++ {
        start, end, val := queries[i][0], queries[i][1], queries[i][2]
        differenceArray[start] += val
        if end + 1 < n {
            differenceArray[end + 1] -= val
        }
    }

    current := 0
    for i := 0; i < n; i++ {
        current += differenceArray[i]
        if nums[i] > current { // Can't reduce this element to ≤ 0
            return false
        }
    }
    return true
}