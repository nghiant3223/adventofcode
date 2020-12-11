package day2

import (
	"adventofcode/2019/fileutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgramAlarm1202(t *testing.T) {
	for _, tc := range []struct {
		name  string
		in    []int
		outIn []int
		out   int
	}{
		{
			name:  "1",
			in:    []int{1, 0, 0, 0, 99},
			outIn: []int{2, 0, 0, 0, 99},
			out:   2,
		},
		{
			name:  "2",
			in:    []int{2, 3, 0, 3, 99},
			outIn: []int{2, 3, 0, 6, 99},
			out:   2,
		},
		{
			name:  "3",
			in:    []int{2, 4, 4, 5, 99, 0},
			outIn: []int{2, 4, 4, 5, 99, 9801},
			out:   2,
		},
		{
			name:  "4",
			in:    []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			outIn: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
			out:   30,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := programAlarm1202(tc.in)
			assert.Equal(t, tc.out, out)
			assert.Equal(t, tc.outIn, tc.in)
		})
	}
}

func TestProgramAlarm1202FromFile(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 6327510

	codes, err := fileutil.ReadIntSlice(inputFile)
	assert.NoError(t, err)
	codes[1], codes[2] = 12, 2
	ans := programAlarm1202(codes)
	assert.Equal(t, ans, expectedOut)
}
