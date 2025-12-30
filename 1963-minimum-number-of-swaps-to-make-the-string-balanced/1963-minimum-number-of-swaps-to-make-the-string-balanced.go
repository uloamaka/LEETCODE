func minSwaps(s string) int {
    balance := 0 // current unmatched '[' count
    bad := 0     // number of unmatched ']'

    for _, ch := range s {
        if ch == '[' {
            balance++
        } else { // ch == ']'
            if balance > 0 {
                balance--
            } else {
                bad++
            }
        }
    }

    // Each swap fixes two bad closings
    return (bad + 1) / 2
}
