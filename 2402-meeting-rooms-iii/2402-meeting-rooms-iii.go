type Room struct {
	id        int
	available int
}

type OccupiedHeap []Room
type FreeHeap []int

func (h OccupiedHeap) Len() int { return len(h) }
func (h OccupiedHeap) Less(i, j int) bool {
	if h[i].available == h[j].available {
		return h[i].id < h[j].id
	}
	return h[i].available < h[j].available
}
func (h OccupiedHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *OccupiedHeap) Push(x interface{}) {
	*h = append(*h, x.(Room))
}
func (h *OccupiedHeap) Pop() interface{} {
	old := *h
	n := len(old)
	room := old[n-1]
	*h = old[:n-1]
	return room
}

func (h FreeHeap) Len() int           { return len(h) }
func (h FreeHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h FreeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *FreeHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *FreeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	roomID := old[n-1]
	*h = old[:n-1]
	return roomID
}

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	meetingCount := make([]int, n)
	freeRooms := &FreeHeap{}
	occupiedRooms := &OccupiedHeap{}

	for i := 0; i < n; i++ {
		heap.Push(freeRooms, i)
	}

	for _, meeting := range meetings {
		start, end := meeting[0], meeting[1]
		duration := end - start

		// Free up rooms available by start time
		for occupiedRooms.Len() > 0 && (*occupiedRooms)[0].available <= start {
			room := heap.Pop(occupiedRooms).(Room)
			heap.Push(freeRooms, room.id)
		}

		if freeRooms.Len() > 0 {
			roomID := heap.Pop(freeRooms).(int)
			heap.Push(occupiedRooms, Room{id: roomID, available: end})
			meetingCount[roomID]++
		} else {
			// All rooms busy, wait for the earliest available room
			room := heap.Pop(occupiedRooms).(Room)
			newEnd := room.available + duration
			heap.Push(occupiedRooms, Room{id: room.id, available: newEnd})
			meetingCount[room.id]++
		}
	}

	// Find room with max meetings (break tie by smaller index)
	maxCount, result := -1, -1
	for i := 0; i < n; i++ {
		if meetingCount[i] > maxCount || (meetingCount[i] == maxCount && i < result) {
			maxCount = meetingCount[i]
			result = i
		}
	}

	return result
}