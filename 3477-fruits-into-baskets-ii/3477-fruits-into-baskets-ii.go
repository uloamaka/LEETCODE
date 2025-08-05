func numOfUnplacedFruits(fruits []int, baskets []int) int {
    used := make([]bool, len(baskets))
    notContained := 0

    for _, fruit := range fruits {
        placed := false

        for i, basket := range baskets {
            if !used[i] && fruit <= basket {
                used[i] = true
                placed = true
                break
            }
        }

        if !placed {
            notContained++
        }
    }

    return notContained
}
