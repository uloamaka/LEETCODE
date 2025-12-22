// func minimumDeletions(s string) int {
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

// func minimumDeletions(s string) int {
//     bCount := 0      // acts like "stack size" of 'b'
//     deletions := 0   // min deletions so far

//     for _, ch := range s {
//         if ch == 'b' {
//             // push 'b' onto the virtual stack
//             bCount++
//         } else { // ch == 'a'
//             // violation if we keep this 'a'
//             // choose cheaper option:
//             // 1) delete this 'a'        -> deletions + 1
//             // 2) delete all previous b  -> bCount
//             if deletions+1 < bCount {
//                 deletions = deletions + 1
//             } else {
//                 deletions = bCount
//             }
//         }
//     }

//     return deletions
// }

func minimumDeletions(s string) int {
    n := len(s)
    dp := make([]int, n)

    bCount := 0

    for i := 0; i < n; i++ {
        if s[i] == 'b' {
            if i > 0 {
                dp[i] = dp[i-1]
            }
            bCount++
        } else { // s[i] == 'a'
            if i == 0 {
                dp[i] = 0
            } else {
                // min(dp[i-1] + 1, bCount)
                if dp[i-1]+1 < bCount {
                    dp[i] = dp[i-1] + 1
                } else {
                    dp[i] = bCount
                }
            }
        }
    }

    return dp[n-1]
}
