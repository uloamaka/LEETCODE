func matchPlayersAndTrainers(players []int, trainers []int) int {
    sort.Ints(players)
    sort.Ints(trainers)

    i := len(players) - 1
    j := len(trainers) - 1
    count := 0

    for i >= 0 && j >= 0 {
        if players[i] <= trainers[j] {
            count++
            j-- 
            i-- 
        } else {
            i-- 
        }
    }

    return count
}