func findArray(pref []int) []int {
    arr := make([]int, len(pref))
    arr[0] = pref[0]
    
    for i := 1; i < len(pref); i++ {
        arr[i] = pref[i-1] ^ pref[i] 
    }
    return arr
}

// initialise a new array of int with the length of the pref array
// loop through the pref array, 
// the 1st index remains the same.
// then apply this arr[i] = pref[i] = pref[i-1] 