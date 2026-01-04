func reverseParentheses(s string) string {
    stack := []int{}
    var str []string

    for _, rn := range s {
        if rn == '(' {
            stack = append(stack, len(str))
            continue
        }
        if rn == ')' {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            rvsd := Reverse(str[top:])
            str = append(str[:top], rvsd...)
            continue
        }
        str = append(str, string(rn))
    }
    return Join(str)
}

func Reverse(s []string) []string {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
    return s
}


func Join(slice []string) string {
    var b strings.Builder
    for _, s := range slice {
        b.WriteString(s)
    }
    return b.String()
}

