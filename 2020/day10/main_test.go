package day10

import (
	"adventofcode/2020/futil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []int
		out  int
	}{
		{
			name: "1",
			in:   []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4},
			out:  35,
		},
		{
			name: "2",
			in:   []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3},
			out:  220,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := adapterArray(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestFromFile(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 2592

	codes, err := futil.ReadIntSlice(inputFile)
	assert.NoError(t, err)
	ans := adapterArray(codes)
	assert.Equal(t, ans, expectedOut)
}

func TestTablePart2(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []int
		out  int
	}{
		{
			name: "1",
			in:   []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4},
			out:  8,
		},
		{
			name: "2",
			in:   []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3},
			out:  19208,
		},
		{
			name: "3",
			in:   []int{10, 6, 4, 7, 1, 5},
			out:  4,
		},
		{
			name: "4",
			in:   []int{17, 6, 10, 5, 13, 7, 1, 4, 12, 11, 14},
			out:  28,
		},
		{
			name: "5",
			in:   []int{4, 11, 7, 8, 1, 6, 5},
			out:  7,
		},
		{
			name: "6",
			in:   []int{1, 2},
			out:  2,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := adapterArrayPart2(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestFromFilePart2(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 198428693313536

	codes, err := futil.ReadIntSlice(inputFile)
	assert.NoError(t, err)
	ans := adapterArrayPart2(codes)
	assert.Equal(t, expectedOut, ans)
}
