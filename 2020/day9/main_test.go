package day9

import (
	"adventofcode/2020/futil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTablePart1(t *testing.T) {
	for _, tc := range []struct {
		name string
		in1  []int
		in2  int
		out  int
	}{
		{
			name: "1",
			in1:  []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576},
			in2:  5,
			out:  127,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := part1(tc.in1, tc.in2)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestFromFilePart1(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 88311122

	size := 25
	codes, err := futil.ReadIntSlice(inputFile)
	assert.NoError(t, err)
	out := part1(codes, size)
	assert.Equal(t, expectedOut, out)
}

func TestTablePart2(t *testing.T) {
	for _, tc := range []struct {
		name string
		in1  []int
		in2  int
		out  int
	}{
		{
			name: "1",
			in1:  []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576},
			in2:  127,
			out:  62,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := part2(tc.in1, tc.in2)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestFromFilePart2(t *testing.T) {
	inputFile := "./input.txt"
	target := 88311122
	expectedOut := 13549369

	codes, err := futil.ReadIntSlice(inputFile)
	assert.NoError(t, err)
	out := part2(codes, target)
	assert.Equal(t, expectedOut, out)
}
