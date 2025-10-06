func clearDigits(s string) string {
    stack := []byte{}
    
    for i:= 0; i < len(s);  i++ {
        if s[i] < '0' || s[i] > '9' {
            stack = append(stack, s[i])
        } else {
            stack = stack[:len(stack)-1]
        }
    }
    return string(stack)
}