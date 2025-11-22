func countPalindromicSubsequence(s string) int {
    result := 0
    azSlc := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
    for _, char := range azSlc {
        firstIndex := strings.Index(s, char)
        lastIndex := strings.LastIndex(s, char)
        if lastIndex > firstIndex+1 {
            cnt := findUniqueCharsCnt(s[firstIndex:lastIndex])
            result += cnt
        }
    }
    return result
}

func findUniqueCharsCnt(s string) int {
	charSet := make(map[rune]bool)
	var cnt int

    if len(s) <= 1 {
        return 1
    }

	for _, char := range s[1:] {
		if !charSet[char] {
			charSet[char] = true
            cnt += 1
		}
	}
	return cnt
}