func maxDifference(s string) int { 
    xMaps := make(map[rune]int)
    for _, xter := range s {
		xMaps[xter]++ 
	}

    // n := len(xMaps)
   	sm := rankByXterCount(xMaps)

	var (
		maxEven, minEven int
		maxOdd, minOdd   int
		foundEven        bool
		foundOdd         bool
	)

	// Initialize minEven/minOdd with max possible, maxEven/maxOdd with min possible
	minEven = 1 << 31    // Start with a very large number
	minOdd = 1 << 31     // Start with a very large number
	maxEven = -1 << 31   // Start with a very small number
	maxOdd = -1 << 31    // Start with a very small number

	for _, num := range sm {
		if checkEven(num) {
			foundEven = true
			if num > maxEven {
				maxEven = num
			}
			if num < minEven {
				minEven = num
			}
		} else {
			foundOdd = true
			if num > maxOdd {
				maxOdd = num
			}
			if num < minOdd {
				minOdd = num
			}
		}
	}

	if !foundEven || !foundOdd {
		return 0 // No even or odd frequencies
	}

	// The maximum difference can be either:
	// (maxEven - minOdd) or (maxOdd - minEven)
	diff1 := maxEven - minOdd
	diff2 := maxOdd - minEven

	if diff1 > diff2 {
		return diff1
	}
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