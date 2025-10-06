// func smallestNumber(pattern string) string {
//     num := []int{}
//     // maxVal := len(pattern) + 1
//     for i := 1; i <= len(pattern)+1; i++ {
//         num = append(num, i)
//     }
//     fmt.Println(num)
//     for i, char := range pattern {
//         if char == 'I'
        
//     }
//     fmt.Println(num)
//     return ""
// }
// 0-1-2-3-4-5-6-7                 
// I-I-I-D-I-D-D-D

// "0,1,2,3,4,5,6,7,8" INDEX
// "1,2,3,4,5,6,7,8,9" NUM
// TO
// "1,2,3,5,4,9,8,7,6"
func smallestNumber(pattern string) string {
    result := ""
    stack := []int{}
    for i := 0; i <= len(pattern); i++ {
        stack = append(stack, i+1)

        if i == len(pattern) || pattern[i] == 'I'{
            for len(stack) > 0 {
                top := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                result += strconv.Itoa(top)
            }
        }
    }
    return result
}

// func popAllOut(stack []int) []int {
//     store := []int{0, len(stack)}
//     for len(stack) > 0 {
//         top := stack[len(stack)-1]
//         stack = stack[:len(stack)-1]
//         store = append(store, top)
//     }
//     return store
// }