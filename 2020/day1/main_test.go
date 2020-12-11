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
	ans := reportRepair(codes)
	assert.Equal(t, ans, expectedOut)
}
