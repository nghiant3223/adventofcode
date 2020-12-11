package day3

import (
	"adventofcode/2020/futil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []string
		out  int
	}{
		{
			name: "1",
			in: []string{
				"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#",
			},
			out: 7,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := numberOfTreeEncounter(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestInputFile(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 211

	codes, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	ans := numberOfTreeEncounter(codes)
	assert.Equal(t, ans, expectedOut)
}

func TestTablePart2(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []string
		out  int
	}{
		{
			name: "1",
			in: []string{
				"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#",
			},
			out: 336,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := numberOfTreeEncounterPart2(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestInputFilePart2(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 3584591857

	codes, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	ans := numberOfTreeEncounterPart2(codes)
	assert.Equal(t, ans, expectedOut)
}
