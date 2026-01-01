func plusOne(digits []int) []int {
    if digits[len(digits)-1] < 9 {
        digits[len(digits)-1] = digits[len(digits)-1] +1
        return digits
    }

    return plusOneRecursive(digits, len(digits)-1)
}

func plusOneRecursive(digits []int, i int) []int {
    if i < 0 {
        // All digits were 9, need a new digit at front
        return append([]int{1}, digits...)
    }

    if digits[i] < 9 {
        digits[i] += 1
        return digits
    }

    digits[i] = 0
    return plusOneRecursive(digits, i-1)
}
