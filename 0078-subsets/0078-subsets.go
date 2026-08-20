func subsets(nums []int) [][]int {
    result := [][]int{}

    var explore func(arr []int, i int, temp []int)

    explore = func(arr []int, i int, temp []int) {
        if len(arr) == i {
            copyTemp := append([]int{}, temp...)
            result = append(result, copyTemp)
            return
        }

        temp = append(temp, arr[i])
        explore(arr, i+1, temp)

        temp = temp[:len(temp)-1]
        explore(arr, i+1, temp)
    }

    explore(nums, 0, []int{})

    return result
}