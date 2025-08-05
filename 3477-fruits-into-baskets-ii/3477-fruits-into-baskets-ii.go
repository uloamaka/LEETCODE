func numOfUnplacedFruits(fruits []int, baskets []int) int {
    notContained := 0
    
    for _, fruit := range fruits {
        placed := false
        
        for i, basket := range baskets {
            if fruit <= basket {
                baskets = append(baskets[:i], baskets[i+1:]...)
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