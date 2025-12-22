// func minimumDeletions(s string) int {
//     // stack := []rune{}
//     // del_a := 0
//     // del_b := 0
//     // for _, rn := range s {
//     //     if len(stack) == 0  {
//     //         stack = append(stack, rn)
//     //     } 
//     //     if stack[len(stack)-1] == 'b' && rn == 'a' {
//     //         del_a++
//     //         continue
//     //     } 
//     //     stack = append(stack,rn)
//     // }

//     // // for _, rn := range s {
//     // //     if len(stack) == 0 {
//     // //         stack = append(stack, rn)
//     // //     }
//     // //     if stack[len(stack)-1] == 'b' && rn == 'a' {
//     // //         del_a++
//     // //         continue
//     // //     } 
//     // //     stack = append(stack,rn)
//     // // }
//     // // return min(del_a, del_b)
//     // return del_a

//      // Step 1: count total 'a'
//     aRight := 0
//     for _, ch := range s {
//         if ch == 'a' {
//             aRight++
//         }
//     }

//     // Step 2: scan and compute minimal cost
//     bLeft := 0
//     minDel := aRight // split before the string

//     for _, ch := range s {
//         if ch == 'a' {
//             aRight--
//         } else { // ch == 'b'
//             bLeft++
//         }

//         // cost at this split
//         if bLeft+aRight < minDel {
//             minDel = bLeft + aRight
//         }
//     }

//     return minDel
// }

func minimumDeletions(s string) int {
    bCount := 0      // acts like "stack size" of 'b'
    deletions := 0   // min deletions so far

    for _, ch := range s {
        if ch == 'b' {
            // push 'b' onto the virtual stack
            bCount++
        } else { // ch == 'a'
            // violation if we keep this 'a'
            // choose cheaper option:
            // 1) delete this 'a'        -> deletions + 1
            // 2) delete all previous b  -> bCount
            if deletions+1 < bCount {
                deletions = deletions + 1
            } else {
                deletions = bCount
            }
        }
    }

    return deletions
}