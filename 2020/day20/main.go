package day20

func part1(lines []string) int {
	tiles := prepareInput(lines)
	tilesMatches := buildTilesMatches(tiles)
	tilesUniqueMatches := buildUniqueMatch(tilesMatches)

	answer := 1
	for id, match := range tilesUniqueMatches {
		if len(match) == 2 {
			answer *= id
		}
	}
	return answer
}

func buildUniqueMatch(tilesMatches map[int][]Match) map[int][]UniqueMatch {
	tilesUniqueMatches := make(map[int][]UniqueMatch)
	for id, matches := range tilesMatches {
		uniqueMatchesSet := make(map[UniqueMatch]struct{})
		for _, match := range matches {
			uniqueMatch := UniqueMatch{
				SourceID: match.SourceID,
				DestID:   match.DestID,
			}
			uniqueMatchesSet[uniqueMatch] = struct{}{}
		}
		i := 0
		uniqueMatches := make([]UniqueMatch, len(uniqueMatchesSet))
		for uniqueMatch := range uniqueMatchesSet {
			uniqueMatches[i] = uniqueMatch
			i++
		}
		tilesUniqueMatches[id] = uniqueMatches
	}
	return tilesUniqueMatches
}

func buildTilesMatches(tiles map[int]*Tile) map[int][]Match {
	allMatches := make(map[int][]Match)
	for _, tile := range tiles {
		tileMatchesSet := make(map[Match]struct{})
		for _, otherTile := range tiles {
			if tile == otherTile {
				continue
			}
			for _, matchFn := range matchFns {
				matches := matchTwoTiles(tile, otherTile, matchFn)
				for _, match := range matches {
					tileMatchesSet[match] = struct{}{}
				}
			}
		}
		i := 0
		tileMatches := make([]Match, len(tileMatchesSet))
		for match := range tileMatchesSet {
			tileMatches[i] = match
			i++
		}
		allMatches[tile.ID] = tileMatches
	}
	return allMatches
}

type matchFn func(tile, otherTile *Tile) bool

var matchFns = []matchFn{
	func(tile, otherTile *Tile) bool {
		return tile.BorderRight == otherTile.BorderLeft
	},
	func(tile, otherTile *Tile) bool {
		return tile.BorderBottom == otherTile.BorderTop
	},
	func(tile, otherTile *Tile) bool {
		return tile.BorderLeft == otherTile.BorderRight
	},
	func(tile, otherTile *Tile) bool {
		return tile.BorderLeft == otherTile.BorderRight
	},
}

func matchTwoTiles(tile, otherTile *Tile, matchFn matchFn) []Match {
	var matches []Match
	for i := Rotate0Deg; i < EndOfTransition; i++ {
		for j := Rotate0Deg; j < EndOfTransition; j++ {
			transitTile := tile.Transposition(i)
			transitOtherTile := otherTile.Transposition(j)
			if matchFn(transitTile, transitOtherTile) {
				match := Match{
					SourceID:    tile.ID,
					SourceState: i,
					DestID:      otherTile.ID,
					DestState:   j,
				}
				matches = append(matches, match)
			}
		}
	}
	return matches
}
