package day11

import (
	"adventofcode/2020/futil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   [][]byte
		out  int
	}{
		{
			name: "1",
			in: [][]byte{
				[]byte("L.LL.LL.LL"),
				[]byte("LLLLLLL.LL"),
				[]byte("L.L.L..L.."),
				[]byte("LLLL.LL.LL"),
				[]byte("L.LL.LL.LL"),
				[]byte("L.LLLLL.LL"),
				[]byte("..L.L....."),
				[]byte("LLLLLLLLLL"),
				[]byte("L.LLLLLL.L"),
				[]byte("L.LLLLL.LL"),
			},
			out: 37,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := part1(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}


func TestFromInputPart(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 3173135507987

	lines, err := futil.ReadBytesSlice(inputFile)
	assert.NoError(t, err)
	out := part1(lines)
	assert.Equal(t, expectedOut, out)
}
