func findLucky(arr []int) int {
    freqMp := make(map[int]int)
    for _, elem := range arr  {
        freqMp[elem]++
    }

    maxLucky := -1
    for key, freq := range freqMp {
        if key == freq && key > maxLucky {
            maxLucky = key
        }
    }

    return maxLucky
}
// use a freq map
//  get the length of the map
//  iterate through the map to get the key that is equal to the cnt