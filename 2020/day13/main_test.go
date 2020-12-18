package day10

import (
	"adventofcode/2020/futil"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	for _, tc := range []struct {
		name string
		in2  []int
		in1  int
		out  int
	}{
		{
			name: "1",
			in1:  939,
			in2:  []int{7, 13, 59, 31, 19},
			out:  295,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := shuttleSearch(tc.in1, tc.in2)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestFromFile(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 174

	lines, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	departmentID, err := strconv.Atoi(lines[0])
	assert.NoError(t, err)
	busStrings := strings.Split(lines[1], ",")
	buses := make([]int, 0, len(busStrings))
	for _, busString := range busStrings {
		if busString == "x" {
			continue
		}
		bus, err := strconv.Atoi(busString)
		assert.NoError(t, err)
		buses = append(buses, bus)
	}
	out := shuttleSearch(departmentID, buses)
	assert.Equal(t, expectedOut, out)
}
