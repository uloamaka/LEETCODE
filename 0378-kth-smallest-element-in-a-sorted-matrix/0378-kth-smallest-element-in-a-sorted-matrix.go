func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    left := matrix[0][0]
    right := matrix[n-1][n-1]

    for left < right {
        mid := left + (right - left) / 2
        count := countElementLessThanOrEqual(matrix, mid)

        if count >= k {
            right = mid 
        } else {
            left = mid + 1 
        }
    }
    return right
}

func countElementLessThanOrEqual(matrix [][]int, target int) int {
    n := len(matrix)
    count := 0
    row := n - 1
    col := 0

    for row >= 0 && col < n {
        if matrix[row][col] <= target {
            count += (row + 1) 
            col++ 
        } else {
            row-- 
        }  
    }
    return count
}