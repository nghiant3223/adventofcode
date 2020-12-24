package day11

const (
	empty    = 'L'
	occupied = '#'
)

func part1(grid [][]byte) int {
	target, changeCount := applyRules(grid)
	for changeCount != 0 {
		target, changeCount = applyRules(target)
	}
	return countOccupiedSeat(target)
}

func applyRules(initial [][]byte) ([][]byte, int) {
	changeCount := 0
	result := make([][]byte, len(initial))
	for i := range result {
		result[i] = make([]byte, len(initial[i]))
	}
	for i, line := range initial {
		for j := range line {
			adjacentOccupiedSeatCount := countOccupiedAdjacentSeat(initial, i, j)
			if initial[i][j] == empty && adjacentOccupiedSeatCount == 0 {
				result[i][j] = occupied
				changeCount++
			} else if initial[i][j] == occupied && adjacentOccupiedSeatCount >= 4 {
				result[i][j] = empty
				changeCount++
			} else {
				result[i][j] = initial[i][j]
			}
		}
	}
	return result, changeCount
}

func countOccupiedAdjacentSeat(grid [][]byte, i, j int) int {
	count := 0
	height := len(grid)
	width := len(grid[0])
	for ii := i - 1; ii <= i+1; ii++ {
		for jj := j - 1; jj <= j+1; jj++ {
			if i == ii && j == jj {
				continue
			}
			if ii < 0 || ii >= height || jj < 0 || jj >= width {
				continue
			}
			if grid[ii][jj] == occupied {
				count++
			}
		}
	}
	return count
}

func countOccupiedSeat(grid [][]byte) int {
	count := 0
	for i, line := range grid {
		for j := range line {
			if grid[i][j] == occupied {
				count++
			}
		}
	}
	return count
}
