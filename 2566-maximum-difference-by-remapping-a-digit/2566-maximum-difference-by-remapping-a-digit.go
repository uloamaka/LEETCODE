func minMaxDifference(num int) int {
    strNumber := strconv.Itoa(num)
    var digitToMax, digitToMin rune
    
    for _, r := range strNumber {
        if digitToMax == 0 && r != '9' {
            digitToMax = r
        }
        if digitToMin == 0 && r != '0' {
            digitToMin = r
        }
        if digitToMax != 0 && digitToMin != 0 {
            break
        }
    }
    
    if digitToMax == 0 {
        digitToMax = '9'
    }
    if digitToMin == 0 {
        digitToMin = '0'
    }
    
    var maxStr, minStr strings.Builder
    for _, r := range strNumber {
        
        if r == digitToMax {
            maxStr.WriteRune('9')
        } else {
            maxStr.WriteRune(r)
        }

        if r == digitToMin {
            minStr.WriteRune('0')
        } else {
            minStr.WriteRune(r)
        }
    }
    
    maxNum, _ := strconv.Atoi(maxStr.String())
    minNum, _ := strconv.Atoi(minStr.String())
    return maxNum - minNum
}