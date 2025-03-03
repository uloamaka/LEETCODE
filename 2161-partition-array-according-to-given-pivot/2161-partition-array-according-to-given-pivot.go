func pivotArray(nums []int, pivot int) []int {
    ans := make([]int, 0 , len(nums))
    m := make(map[string][]int)
    for _, num := range nums {
    if num < pivot {
        m["lessThan"] = append(m["lessThan"], num)
    } else if num == pivot {
        m["equalTo"] = append(m["equalTo"], num)
    } else {
        m["greaterThan"] = append(m["greaterThan"], num)
    }
    }
    fmt.Println(m["lessThan"])
    ans = append(ans, m["lessThan"]...)
    ans = append(ans, m["equalTo"]...)
    ans = append(ans, m["greaterThan"]...)
    return ans 
}