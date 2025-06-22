func divideString(s string, k int, fill byte) []string {
    sliceSize := len(s)
    strSlice := []string{}
    subSlice := sliceSize/k
    start, end := 0, k
    for i := 0; i < subSlice; i++ {
        strSlice = append(strSlice, string(s[start:end]))
        start += subSlice
        end += subSlice
    }
    // use the fill for the last group
    mod := sliceSize%k
    if mod != 0 {
        extraIdx := k - mod 
        extraCharIdx := sliceSize - mod
        str := s[extraCharIdx:]
        for i := 0; i < extraIdx; i++ {
            str += string(fill)
        }
        strSlice = append(strSlice, string(str))
    }
   
    return strSlice
}