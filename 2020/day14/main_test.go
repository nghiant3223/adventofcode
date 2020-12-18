package day14

import (
	"adventofcode/2020/futil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Table(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []string
		out  int64
	}{
		{
			name: "1",
			in: []string{
				"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
				"mem[8] = 11",
				"mem[7] = 101",
				"mem[8] = 0",
			},
			out: 165,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := part1(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestFromFilePart1(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 13865835758282

	codes, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	ans := part1(codes)
	assert.Equal(t, expectedOut, ans)
}

func TestGenerateCombinations(t *testing.T) {
	for _, tc := range []struct {
		name string
		in1  string
		in2  int
		out  []string
	}{
		{
			name: "1",
			in1:  "X0X",
			in2:  3,
			out: []string{
				"000",
				"001",
				"100",
				"101",
			},
		},
		{
			name: "1",
			in1:  "XXX",
			in2:  3,
			out: []string{
				"000",
				"001",
				"010",
				"011",
				"100",
				"101",
				"110",
				"111",
			},
		},
		{
			name: "1",
			in1:  "00X",
			in2:  3,
			out: []string{
				"000",
				"001",
			},
		},
		{
			name: "1",
			in1:  "X00",
			in2:  3,
			out: []string{
				"000",
				"100",
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := generateCombinations(tc.in2, tc.in1)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestPart2Table(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []string
		out  int64
	}{
		{
			name: "1",
			in: []string{
				"mask = 000000000000000000000000000000X1001X",
				"mem[42] = 100",
				"mask = 00000000000000000000000000000000X0XX",
				"mem[26] = 1",
			},
			out: 208,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := part2(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestFromFilePart2(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 924192532172

	codes, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	ans := part2(codes)
	assert.Equal(t, expectedOut, ans)
}
