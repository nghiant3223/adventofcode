package day21

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
				"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
				"trh fvjkl sbzzf mxmxvkd (contains dairy)",
				"sqjhc fvjkl (contains soy)",
				"sqjhc mxmxvkd sbzzf (contains fish)",
			},
			out: 5,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := part1(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestPart1FromInput(t *testing.T) {
		inputFile := "./input.txt"
		expectedOut := 187

		lines, err := futil.ReadStringSlice(inputFile)
		assert.NoError(t, err)
		out := part1(lines)
		assert.Equal(t, expectedOut, out)
}
