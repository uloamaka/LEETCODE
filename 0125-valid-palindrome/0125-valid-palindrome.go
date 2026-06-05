func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    l := 0
    r := len(s)-1

    for l < r {
        for l < r {
            if !isAlphaNumeric(s[l]) {
                l++
            }else {
                break
            }
        }
        for r > l {
            if !isAlphaNumeric(s[r]) {
                r--
            } else {
                break
            }
        }

        if s[l] == s[r] {
            l++
            r--
        } else {
            return false
        }  
    }
    return true
}

func isAlphaNumeric(c byte) bool {
 if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
  return true
 }

 return false
}