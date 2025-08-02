func minCost(basket1 []int, basket2 []int) int64 {
    count := make(map[int]int)
    for _, fruit1 := range basket1 {
        count[fruit1]++
    }
    for _, fruit2 := range basket2 {
        count[fruit2]--
    }

    surplus := []int{}
    for val, freq := range count {
        if freq % 2 != 0 {
            return -1
        } 

        for i := 0; i < abs(freq)/2; i++ {
            if freq > 0 {
                surplus = append(surplus, val)
            } else {
                surplus = append(surplus, val)
            }
        }
    }

    if len(surplus) == 0 {
		return 0
	}

	sort.Ints(surplus)

    all := append(basket1, basket2...)
    minElm := all[0]
    for _, v := range all {
        if v < minElm {
            minElm = v
        }
    }

    n := len(surplus)
    cost := 0
    for i := 0; i <n/2; i++ {
        cost += min(surplus[i], 2*minElm)
    }

    return int64(cost)
}

func abs(num int)int {
    if num < 0 {
        return -num
    }
    return num
}

// Counts fruit frequencies in both baskets.
// Checks feasibility: If any fruit has an odd total count, return -1 (can't balance).
// Calculates target count for each fruit (half of total).
// Identifies excess fruits in each basket that need to be swapped.
// Sorts excesses:
// excess1 ascending (cheapest first),
// excess2 descending (most expensive first).
// Computes swap cost:
// Either swap directly (min(excess1[i], excess2[i])),
// Or do a double swap via the cheapest fruit (2 * minimumCost).
// Returns total minimum cost.