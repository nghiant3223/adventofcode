package day20

import (
	"regexp"
	"strconv"
)

var headerPattern = regexp.MustCompile(`Tile (\d+):`)

func prepareInput(lines []string) map[int]*Tile {
	var tileLines [][]string
	var currentTileLines []string

	for _, line := range lines {
		if line == "" {
			tileLines = append(tileLines, currentTileLines)
			currentTileLines = []string{}
			continue
		}
		currentTileLines = append(currentTileLines, line)
	}
	if len(currentTileLines) != 0 {
		tileLines = append(tileLines, currentTileLines)
	}

	lookup := make(map[int]*Tile, len(tileLines))
	for _, lines := range tileLines {
		tile := &Tile{}
		match := headerPattern.FindStringSubmatch(lines[0])
		tile.ID, _ = strconv.Atoi(match[1])
		height := len(lines) - 1
		width := len(lines[0])
		tile.BorderTop = lines[1]
		tile.BorderBottom = lines[height]
		for i := 1; i <= height; i++ {
			line := lines[i]
			tile.BorderLeft += string(line[0])
			tile.BorderRight += string(line[width-1])
		}
		lookup[tile.ID] = tile
	}
	return lookup
}

