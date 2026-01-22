func clearDigits(s string) string {
    stack := make([]byte, 0)
    
    for i:= 0; i < len(s);  i++ {
        if s[i] < '0' || s[i] > '9' {
            stack = append(stack, s[i])

            continue
        } 
        stack = stack[:len(stack)-1]
    }
    return string(stack)
}