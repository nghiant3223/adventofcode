package day2

import (
	"adventofcode/2020/futil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []string
		out  int
	}{
		{
			name: "1",
			in: []string{
				"1-3 a: abcde",
				"1-3 b: cdefg",
				"2-9 c: ccccccccc",
			},
			out: 2,
		},
		{
			name: "2",
			in: []string{
				"3-4 c: crlcvcxwd",
				"3-9 s: vjbsxhwwdvshfxstc",
				"7-8 f: jfhnffbgsfjfwf",
				"11-12 c: cccccccccccc",
			},
			out: 3,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out, err := numberOfValidPasswords(tc.in)
			assert.NoError(t, err)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestInputFile(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 546

	codes, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	ans, err := numberOfValidPasswords(codes)
	assert.NoError(t, err)
	assert.Equal(t, ans, expectedOut)
}

func TestInputFilePart2(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 275

	codes, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	ans, err := numberOfValidPasswordsPart2(codes)
	assert.NoError(t, err)
	assert.Equal(t, ans, expectedOut)
}
