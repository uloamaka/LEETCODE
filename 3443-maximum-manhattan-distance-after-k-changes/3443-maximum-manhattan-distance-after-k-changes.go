func maxDistance(s string, k int) int { 
    ans := 0
    north, south, east, west := 0,0,0,0
    count := func(drt1, drt2, times int) int {
        return int(math.Abs(float64(drt1 - drt2))) + times * 2
    }

    for _, it := range s {
        switch it {
            case 'N': 
                north++
            case 'S':
                south++
            case 'E':
                east++
            case 'W':
                west++
        }
        times1 := min(min(north, south), k)
        times2 := min(min(east, west), k - times1)
        current := count(north, south, times1) +count(east, west, times2)
        if current > ans {
            ans = current
        }
    }
    return ans
}