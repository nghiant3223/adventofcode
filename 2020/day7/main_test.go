package day7

import (
	"adventofcode/2020/futil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []string
		out  int
	}{
		{
			name: "1",
			in: []string{
				"light red bags contain 1 bright white bag, 2 muted yellow bags.",
				"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
				"bright white bags contain 1 shiny gold bag.",
				"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
				"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
				"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
				"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
				"faded blue bags contain no other bags.",
				"dotted black bags contain no other bags.",
			},
			out: 4,
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
	expectedOut := 115

	codes, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	ans := part1(codes)
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
				"light red bags contain 1 bright white bag, 2 muted yellow bags.",
				"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
				"bright white bags contain 1 shiny gold bag.",
				"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
				"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
				"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
				"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
				"faded blue bags contain no other bags.",
				"dotted black bags contain no other bags.",
			},
			out: 4,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out := part1(tc.in)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestFromFilePart2(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 115

	codes, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	ans := part2(codes)
	assert.Equal(t, expectedOut, ans)
}
