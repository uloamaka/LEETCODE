func isValid(word string) bool {
    n := len(word)
    if n < 3 {
        return false
    }
    if !strings.ContainsAny(word,"aeiouAEIOU") {
        return false
    }
    if !strings.ContainsAny(word,"bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ") {
        return false
    }
    if strings.ContainsAny(word,"@#$") {
        return false
    }

    return true
}