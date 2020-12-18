package day1

import (
	"adventofcode/2020/futil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   string
		out  int
	}{
		{
			name: "1",
			in:   "FBFBBFFRLR",
			out:  357,
		},
		{
			name: "2",
			in:   "BFFFBBFRRR",
			out:  567,
		},
		{
			name: "3",
			in:   "FFFBBBFRRR",
			out:  119,
		},
		{
			name: "4",
			in:   "BBFFBBFRLL",
			out:  820,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := binaryBoarding(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestReportRepairFromFile(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 826

	codes, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	ans := maxSeadID(codes)
	assert.Equal(t, ans, expectedOut)
}

func TestReportRepairFromFilePart2(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 638

	codes, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	ans := mySeatID(codes)
	assert.Equal(t, ans, expectedOut)
}
