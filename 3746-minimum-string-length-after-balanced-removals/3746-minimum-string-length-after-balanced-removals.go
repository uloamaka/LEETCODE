// func minLengthAfterRemovals(s string) int {
//     stack := []rune{}

//     for _, rn := range s {
//         if len(stack) > 0 && rn == 'b' && stack[len(stack)-1] == 'a' {
//             stack = stack[:len(stack)-1]
//             continue
//         }

//         if len(stack) > 0 && rn == 'a' && stack[len(stack)-1] == 'b' {
//             stack = stack[:len(stack)-1]
//             continue
//         }
//         stack = append(stack, rn)
//     }
//     return len(stack)
// }

func minLengthAfterRemovals(s string) int {
    cnt := 0 
    for _, rn := range s {
        if rn == 'a' {
            cnt++
        } else {
            cnt--
        }
    }
    return abs(cnt)
}

func abs(val int) int {
    if val < 0 {
        return -val
    }
    return val
}
