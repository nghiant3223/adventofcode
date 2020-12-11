package day3

const (
	tree         = '#'
	square       = '.'
	goRightStep  = 3
	goBottomStep = 1
)

func numberOfTreeEncounter(rows []string) int {
	width := len(rows[0])
	rowCount := len(rows)
	treeEncounteredCount := 0
	for j, i := 0, goBottomStep; i < rowCount; i += goBottomStep {
		j = (j + goRightStep) % width
		if rows[i][j] == tree {
			treeEncounteredCount++
		}
	}
	return treeEncounteredCount
}

func numberOfTreeEncounterPart2(rows []string) int {
	answer := 1
	width := len(rows[0])
	rowCount := len(rows)
	slopes := [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, slope := range slopes {
		treeEncounteredCount := 0
		for j, i := 0, slope[1]; i < rowCount; i += slope[1] {
			j = (j + slope[0]) % width
			if rows[i][j] == tree {
				treeEncounteredCount++
			}
		}
		answer *= treeEncounteredCount
	}
	return answer
}
