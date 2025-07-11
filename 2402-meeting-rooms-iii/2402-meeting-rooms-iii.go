type Room struct {
	id        int
	available int
}

type RoomHeap []Room

func (h RoomHeap) Len() int           { return len(h) }
func (h RoomHeap) Less(i, j int) bool {
	if h[i].available == h[j].available {
		return h[i].id < h[j].id
	}
	return h[i].available < h[j].available
}
func (h RoomHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *RoomHeap) Push(x interface{}) { *h = append(*h, x.(Room)) }
func (h *RoomHeap) Pop() interface{} {
	old := *h
	n := len(old)
	room := old[n-1]
	*h = old[:n-1]
	return room
}

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	meetingCount := make([]int, n)

	freeRooms := &RoomHeap{}
	occupiedRooms := &RoomHeap{}

	for i := 0; i < n; i++ {
		heap.Push(freeRooms, Room{id: i, available: 0})
	}

	for _, meeting := range meetings {
		start, end := meeting[0], meeting[1]
		duration := end - start

		// Free up rooms that become available by current start time
		for occupiedRooms.Len() > 0 && (*occupiedRooms)[0].available <= start {
			room := heap.Pop(occupiedRooms).(Room)
			heap.Push(freeRooms, Room{id: room.id, available: room.available})
		}

		if freeRooms.Len() > 0 {
			room := heap.Pop(freeRooms).(Room)
			heap.Push(occupiedRooms, Room{id: room.id, available: end})
			meetingCount[room.id]++
		} else {
			room := heap.Pop(occupiedRooms).(Room)
			newEnd := room.available + duration
			heap.Push(occupiedRooms, Room{id: room.id, available: newEnd})
			meetingCount[room.id]++
		}
	}

	// Get the room with the most meetings (smallest index in case of tie)
	maxMeetings, result := -1, -1
	for i := 0; i < n; i++ {
		if meetingCount[i] > maxMeetings || (meetingCount[i] == maxMeetings && i < result) {
			maxMeetings = meetingCount[i]
			result = i
		}
	}

	return result
}