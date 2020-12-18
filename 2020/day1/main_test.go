package day1

import (
	"testing"

	"adventofcode/2020/futil"

	"github.com/stretchr/testify/assert"
)

func TestReportRepairFromFile(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 918339

	codes, err := futil.ReadIntSlice(inputFile)
	assert.NoError(t, err)
	e1, e2 := reportRepair(codes, target)
	assert.Equal(t, e1*e2, expectedOut)
}

func TestReportRepairFromFilePart2(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 918339

	codes, err := futil.ReadIntSlice(inputFile)
	assert.NoError(t, err)
	target := reportRepairPart2(codes, target)
	assert.Equal(t, target, expectedOut)
}
