package day19

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
			out:  2,
			in: []string{
				`0: 4 1 5`,
				`1: 2 3 | 3 2`,
				`2: 4 4 | 5 5`,
				`3: 4 5 | 5 4`,
				`4: "a"`,
				`5: "b"`,
				``,
				`ababbb`,
				`bababa`,
				`abbbab`,
				`aaabbb`,
				`aaaabbb`,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := part1(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestFromInputPart1(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 187

	lines, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	out := part1(lines)
	assert.Equal(t, expectedOut, out)
}
