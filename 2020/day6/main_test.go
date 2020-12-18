package day6

import (
	"adventofcode/2020/futil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []string
		out  int
	}{
		{
			name: "1",
			in: []string{
				"abc",
				"",
				"a",
				"b",
				"c",
				"",
				"ab",
				"ac",
				"",
				"a",
				"a",
				"a",
				"a",
				"",
				"b",
			},
			out: 11,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := numberOfAnswers(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestFromFile(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 6885

	codes, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	ans := numberOfAnswers(codes)
	assert.Equal(t, expectedOut, ans)
}

func TestPart2(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []string
		out  int
	}{
		{
			name: "1",
			in: []string{
				"abc",
				"",
				"a",
				"b",
				"c",
				"",
				"ab",
				"ac",
				"",
				"a",
				"a",
				"a",
				"a",
				"",
				"b",
			},
			out: 6,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := numberOfAnswersPart2(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestFromFilePart2(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 3550

	codes, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	ans := numberOfAnswersPart2(codes)
	assert.Equal(t, expectedOut, ans)
}
