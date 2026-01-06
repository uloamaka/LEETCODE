func minOperations(logs []string) int {
    cnt := 0

    for _, vl := range logs {
        if vl == "../" {
            if cnt > 0 {
                cnt--
            }
            continue
        }

        if vl == "./" {
            continue
        }

        cnt++
    }

    return cnt
}