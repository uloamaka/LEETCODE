func findLucky(arr []int) int {
    freqMp := make(map[int]int)
    for _, elem := range arr  {
        freqMp[elem]++
    }

    newSlice := make([]int, 0)
    for key, elem := range freqMp {
        if key == elem {
            newSlice = append(newSlice, key)
        }
    }
    
    n := len(newSlice)
    sort.Ints(newSlice)
    if n != 0 {
        return newSlice[n-1]
    }
    // fmt.Println()
    return -1
}
// use a freq map
//  get the length of the map
//  iterate through the map to get the key that is equal to the cnt