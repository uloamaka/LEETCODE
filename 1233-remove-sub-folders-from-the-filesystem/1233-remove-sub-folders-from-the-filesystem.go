func removeSubfolders(folder []string) []string {
	// Sort the folders alphabetically
	sort.Strings(folder)

	// Initialize the result slice and add the first folder
	result := []string{folder[0]}

	// Iterate through each folder and check if it's a sub-folder
	for i := 1; i < len(folder); i++ {
		lastFolder := result[len(result)-1] + "/"

		// Check if the current folder starts with the last added folder path
		if !strings.HasPrefix(folder[i], lastFolder) {
			result = append(result, folder[i])
		}
	}

	return result
}