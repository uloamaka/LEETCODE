func mostBooked(n int, meetings [][]int) int {
    sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
    meeting_count := make([]int, n)
	room_availability_time := make(map[int]int, n)
	for i := 0; i < n; i++ {
		room_availability_time[i] = 0
	}

	for _, meeting := range meetings {
		start, end := meeting[0], meeting[1]
		duration := end - start
		assigned := false

		// Try to assign meeting to a free room
		for i := 0; i < n; i++ {
			if room_availability_time[i] <= start {
				room_availability_time[i] = end
				meeting_count[i]++
				assigned = true
				break
			}
		}

		// If no room was free, assign to the one with earliest availability
		if !assigned {
			earliestTime := int(1e18) // some large number
			chosenRoom := -1
			for i := 0; i < n; i++ {
				if room_availability_time[i] < earliestTime {
					earliestTime = room_availability_time[i]
					chosenRoom = i
				}
			}
			room_availability_time[chosenRoom] += duration
			meeting_count[chosenRoom]++
		}
	}

	// Find the room with the most meetings (prefer smaller index in case of tie)
	maxCount := -1
	resultRoom := -1
	for i := 0; i < n; i++ {
		if meeting_count[i] > maxCount || (meeting_count[i] == maxCount && i < resultRoom) {
			maxCount = meeting_count[i]
			resultRoom = i
		}
	}

	return resultRoom
}