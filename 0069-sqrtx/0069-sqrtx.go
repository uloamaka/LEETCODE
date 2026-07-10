func mySqrt(x int) int {
    left, right := 0, x 
    
    if x < 2 {
        return x
    }

    for left < right {
        mid := left + (right-left) / 2

        if mid*mid > x {
            right = mid
        } else {
            left = mid+1
        }
    }
    return left-1
}