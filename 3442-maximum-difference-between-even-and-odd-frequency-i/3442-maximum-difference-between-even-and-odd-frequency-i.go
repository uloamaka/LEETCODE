func maxDifference(s string) int { 
    xMaps := make(map[rune]int)
    for _, xter := range s {
		xMaps[xter]++ 
	}

   	sm := rankByXterCount(xMaps)

	var (
		minEven int
		maxOdd int
		foundEven bool
		foundOdd bool
	)

	minEven = 1 << 31    
	maxOdd = -1 << 31    

	for _, num := range sm {
		if checkEven(num) {
			foundEven = true
			if num < minEven {
				minEven = num
			}
		} else {
			foundOdd = true
			if num > maxOdd {
				maxOdd = num
			}
		}
	}

	diff2 := maxOdd - minEven

	return diff2
}

func rankByXterCount(m map[rune]int) []int {
    var values []int
    for _, v := range m {
        values = append(values, v)
    }

    sort.Slice(values, func(i, j int) bool {
        return values[i] > values[j]  
    })
    return values
}

func checkEven(number int) bool {
	return number%2 == 0
}