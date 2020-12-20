package day8

import (
	"adventofcode/2020/futil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []string
		out  int
	}{
		{
			name: "1",
			in: []string{
				"nop +0",
				"acc +1",
				"jmp +4",
				"acc +3",
				"jmp -3",
				"acc -99",
				"acc +1",
				"jmp -4",
				"acc +6",
			},
			out: 5,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			instructions := prepareInput(tc.in)
			out := part1(instructions)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestFromFilePart1(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 1317

	codes, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	instructions := prepareInput(codes)
	out := part1(instructions)
	assert.Equal(t, expectedOut, out)
}

func TestPart2(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []string
		out  int
	}{
		{
			name: "1",
			in: []string{
				"nop +0",
				"acc +1",
				"jmp +4",
				"acc +3",
				"jmp -3",
				"acc -99",
				"acc +1",
				"jmp -4",
				"acc +6",
			},
			out: 8,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			instructions := prepareInput(tc.in)
			out := part2(instructions)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestFromFilePart2(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 1317

	codes, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	instructions := prepareInput(codes)
	out := part2(instructions)
	assert.Equal(t, expectedOut, out)
}
