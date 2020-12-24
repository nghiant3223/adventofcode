package day16

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
			in: []string{
				"class: 1-3 or 5-7",
				"row: 6-11 or 33-44",
				"seat: 13-40 or 45-50",
				"",
				"your ticket:",
				"7,1,14",
				"",
				"nearby tickets:",
				"7,3,47",
				"40,4,50",
				"55,2,20",
				"38,6,12",
			},
			out: 71,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rules, _, tickets := prepareInput(tc.in)
			out := part1(rules, tickets)
			assert.Equal(t, tc.out, out)
		})
	}
}

func TestFromInput(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 27870

	lines, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	rules, _, tickets := prepareInput(lines)
	out := part1(rules, tickets)
	assert.Equal(t, expectedOut, out)

}

func TestPart2(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []string
		out  []string
	}{
		{
			in: []string{
				"class: 0-1 or 4-19",
				"row: 0-5 or 8-19",
				"seat: 0-13 or 16-19",
				"",
				"your ticket:",
				"11,12,13",
				"",
				"nearby tickets:",
				"3,9,18",
				"15,1,5",
				"5,14,9",
			},
			out: []string{"row", "class", "seat"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rules, _, tickets := prepareInput(tc.in)
			tickets = discardInvalidTickets(rules, tickets)
			out := part2(rules, tickets)
			assert.Equal(t, len(tc.out), len(out))
			for i, rule := range out {
				assert.Equal(t, tc.out[i], rule.Name)
			}
		})
	}
}

func TestFromInputPart2V2(t *testing.T) {
	inputFile := "./input.txt"
	expectedOut := 3173135507987

	lines, err := futil.ReadStringSlice(inputFile)
	assert.NoError(t, err)
	rules, myTicket, tickets := prepareInput(lines)
	tickets = discardInvalidTickets(rules, tickets)
	order := part2(rules, tickets)
	out := getAnswerFromOrder(order, myTicket)
	assert.Equal(t, expectedOut, out)
}
