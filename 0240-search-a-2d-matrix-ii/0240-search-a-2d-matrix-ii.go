func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])

	// Pivot at the top-right corner: this cell is the largest in its row
	// and the smallest in its column, the one spot where the two sort
	// directions point opposite ways and a single compare can prune a line.
	row, col := 0, cols-1

	// Walk while we're still on the board: row can only grow, col can only
	// shrink, so at most rows+cols steps before one pointer falls off.
	for row < rows && col >= 0 {
		currentElement := matrix[row][col]

		if currentElement == target {
			return true
		}

		if target > currentElement {
			// Current is the max of this row, yet still too small — the
			// target cannot be anywhere in this row, so retire the row.
			row += 1
		} else if target < currentElement {
			// Current is the min of this column, yet still too big —
			// nothing below it can match in this column, so retire it.
			col -= 1
		}
	}

	// Pointers crossed off the grid without a hit: target isn't present.
	return false
}