func minSwaps(s string) int {
    stack := []rune{}
    var bad int

    for _, rn := range s {
        if rn == '[' {
            stack = append(stack, rn)
        } 
        if len(stack) > 0 && rn == ']' {
            stack = stack[:len(stack)-1]
            continue
        } 
        if rn == ']' {
            bad++
        }
    }
    return (bad+1)/2
}
