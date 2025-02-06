func findThePrefixCommonArray(A []int, B []int) []int {
    n := len(A)
    result := make([]int, n)
    setA := make(map[int]bool)
    setB := make(map[int]bool)
    
    for i := 0; i < n; i++ {
        count := 0
        setA[A[i]] = true
        setB[B[i]] = true
        
        for num := range setA {
            if setB[num] {
                count++
            }
        }
        
        result[i] = count
    }
    return result
}
// Initialize:
// - result array C of length n
// - setA = empty set to track numbers in A
// - setB = empty set to track numbers in B

// For each index i from 0 to n-1:
//    1. Add A[i] to setA
//    2. Add B[i] to setB
//    3. Initialize count = 0
//    4. For each number in setA:
//       - If number also exists in setB:
//          increment count
//    5. Set C[i] = count