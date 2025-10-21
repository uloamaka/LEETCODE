func maxFrequencyElements(nums []int) int {
    m := make(map[int]int)
    maxC, cnt := 0, 0
    

    for _, num := range nums {
        m[num]++
    }

    //fmt.Println(m)
    
    for _, v := range m {
        if v > maxC {
            maxC = v
            cnt = 1
        } else if v == maxC {
            cnt++
        }
    }

    return cnt*maxC
}