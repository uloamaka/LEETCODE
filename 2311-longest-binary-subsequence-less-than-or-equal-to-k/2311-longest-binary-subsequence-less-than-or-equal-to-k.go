func longestSubsequence(s string, k int) int {
    n := len(s)
    left := n - 1
    right := n
    ans :=  0
    if n == 1 {
        return 1
    }
    i, _ := strconv.ParseInt(s[0:], 2, 64);
    if int64(k) < i {
        for j := n; j > 1; j-- {
            i, _ := strconv.ParseInt(s[left:right], 2, 64); 
            if i <= int64(k) {
                ans++   
                // fmt.Println(ans)     
            } else {
                fmt.Println(ans) 
                break
            }
            left--
        }
        ans += countOnlyZeroFromCurrLeft(left, s) 
        return ans
    } else {
        return n
    }
}

func countOnlyZeroFromCurrLeft(curr_left int, binaryStr string) int {
    cnt := 0
    for _, rn :=  range binaryStr[0:curr_left] {
        if rn == '0' {
            cnt++
        }
    }
    return cnt
}