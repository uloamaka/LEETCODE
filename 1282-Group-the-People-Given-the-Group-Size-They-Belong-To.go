func groupThePeople(groupSizes []int) [][]int {
	sizeToPeople := make(map[int][]int)
	for person, size := range groupSizes {
		sizeToPeople[size] = append(sizeToPeople[size], person)
	}

	result := [][]int{}
	
	for size, people := range sizeToPeople {
		for i := 0; i < len(people); i += size {
			group := people[i:i+size]
			result = append(result, group)
		}
	}
	
	return result
}