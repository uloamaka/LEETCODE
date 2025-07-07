func maxEvents(events [][]int) int {
    // Sort events based on their end day
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	attendedDays := make(map[int]bool)
	count := 0

	for _, event := range events {
		start, end := event[0], event[1]
		for day := start; day <= end; day++ {
			if !attendedDays[day] {
				attendedDays[day] = true
				count++
				break
			}
		}
	}

	return count
}