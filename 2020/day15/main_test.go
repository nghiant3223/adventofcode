package day15

import (
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
			in1:  []int{0, 3, 6},
			in2:  2020,
			out:  436,
		},
		{
			name: "1",
			in1:  []int{1, 3, 2},
			in2:  2020,
			out:  1,
		},
		{
			name: "2",
			in1:  []int{2, 1, 3},
			in2:  2020,
			out:  10,
		},
		{
			name: "3",
			in1:  []int{1, 2, 3},
			in2:  2020,
			out:  27,
		},
		{
			name: "4",
			in1:  []int{2, 3, 1},
			in2:  2020,
			out:  78,
		},
		{
			name: "4",
			in1:  []int{3, 2, 1},
			in2:  2020,
			out:  438,
		},
		{
			name: "4",
			in1:  []int{3, 1, 2},
			in2:  2020,
			out:  1836,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := part1(tc.in1, tc.in2)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestFromInputPart1(t *testing.T) {
	in1 := []int{0, 1, 4, 13, 15, 12, 16}
	in2 := 2020
	expectedOut := 1665
	out := part1(in1, in2)
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
			in1:  []int{0, 3, 6},
			in2:  30000000,
			out:  175594,
		},
		{
			name: "1",
			in1:  []int{1, 3, 2},
			in2:  30000000,
			out:  2578,
		},
		{
			name: "2",
			in1:  []int{2, 1, 3},
			in2:  30000000,
			out:  3544142,
		},
		{
			name: "3",
			in1:  []int{1, 2, 3},
			in2:  30000000,
			out:  261214,
		},
		{
			name: "4",
			in1:  []int{2, 3, 1},
			in2:  30000000,
			out:  6895259,
		},
		{
			name: "4",
			in1:  []int{3, 2, 1},
			in2:  30000000,
			out:  18,
		},
		{
			name: "4",
			in1:  []int{3, 1, 2},
			in2:  30000000,
			out:  362,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := part1(tc.in1, tc.in2)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestFromInputPart2(t *testing.T) {
	in1 := []int{0, 1, 4, 13, 15, 12, 16}
	in2 := 30000000
	expectedOut := 16439
	out := part1(in1, in2)
	assert.Equal(t, expectedOut, out)
}
