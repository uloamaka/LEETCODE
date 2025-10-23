func hasSameDigits(s string) bool {
    n := len(s)
    res := operations(s)
    if n == 2 {
        if s[0] == s[1] {
            return true
        } else {
            return false
        } 
    }
    return hasSameDigits(res)
}

func operations(s string) string{
    var builder strings.Builder
    for i := 0; i < len(s)-1; i++ {
        modular := (s[i]+s[i+1])%10
        builder.WriteByte(modular)
    }  
    
    return builder.String()
}