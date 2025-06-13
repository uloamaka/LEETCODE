func minimumAbsDifference(arr []int) [][]int {
    if sort.IntsAreSorted(arr) != true {
        sort.Ints(arr)
    } 
    
    leftIdx := 0 
    rightIdx := 1
    minSlice := []int{arr[leftIdx], arr[rightIdx]}
    ans := [][]int{minSlice}
    minAbsDiff := abs(arr[leftIdx] - arr[rightIdx])
    for i := 1; i < len(arr)-1; i++ {
        if abs(arr[i] - arr[i+1]) == minAbsDiff {
            ans = append(ans, []int{arr[i],arr[i+1]})
        }
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
