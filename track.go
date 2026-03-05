package main

type Track struct {
	layout [10][15]string
	path   [][2]int // Ordered path coordinates
}

func (track *Track) setLayout(layout [10][15]string) {
	track.layout = layout
	track.parsePath()
}

func (track *Track) getLayout() [10][15]string {
	return track.layout
}

func (track *Track) getPath() [][2]int {
	return track.path
}

// parsePath extracts the path from the layout by following 1s from the start
func (track *Track) parsePath() {
	var path [][2]int

	// Find starting position (1 in top-left area)
	var startRow, startCol int = -1, -1
	for i := 0; i < ROW_COUNT; i++ {
		for j := 0; j < COLUMN_COUNT; j++ {
			if track.layout[i][j] == "1" {
				startRow, startCol = i, j
				break
			}
		}
		if startRow != -1 {
			break
		}
	}

	if startRow == -1 {
		return // No path found
	}

	path = append(path, [2]int{startRow, startCol})

	// Follow the path
	prevRow, prevCol := -1, -1
	currentRow, currentCol := startRow, startCol

	for {
		var found bool = false
		// Check all 4 directions
		directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

		for _, dir := range directions {
			newRow := currentRow + dir[0]
			newCol := currentCol + dir[1]

			// Skip if going back
			if newRow == prevRow && newCol == prevCol {
				continue
			}

			// Check bounds
			if newRow < 0 || newRow >= ROW_COUNT || newCol < 0 || newCol >= COLUMN_COUNT {
				continue
			}

			cell := track.layout[newRow][newCol]
			if cell == "1" || cell == "*" {
				path = append(path, [2]int{newRow, newCol})
				prevRow, prevCol = currentRow, currentCol
				currentRow, currentCol = newRow, newCol
				found = true
				break
			}
		}

		if !found {
			break
		}
	}

	track.path = path
}
