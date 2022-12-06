package day5

import (
	"testing"
)

func TestPart1(t *testing.T) {
	if want, got := "RNZLFZSJH", part1(); want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestPart2(t *testing.T) {
	if want, got := "CNSFCGJSM", part2(); want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

/*
	[D]

[N] [C]
[Z] [M] [P]
1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
*/
func TestMoveCrates2(t *testing.T) {
	testCases := []struct {
		name   string
		moves  []cratemove
		stacks []*Stack
		output string
	}{

		{
			"1",
			[]cratemove{{1, 2, 1}},
			[]*Stack{{[]string{"Z", "N"}}, {[]string{"M", "C", "D"}}, {[]string{"P"}}},
			"DCP",
		},
		{
			"2",
			[]cratemove{{3, 1, 3}},
			[]*Stack{{[]string{"Z", "N", "D"}}, {[]string{"M", "C"}}, {[]string{"P"}}},
			"CD",
		},
		{
			"3",
			[]cratemove{{2, 2, 1}},
			[]*Stack{{[]string{}}, {[]string{"M", "C"}}, {[]string{"P", "Z", "N", "D"}}},
			"CD",
		},
		{
			"4",
			[]cratemove{{1, 1, 2}},
			[]*Stack{{[]string{"M", "C"}}, {[]string{}}, {[]string{"P", "Z", "N", "D"}}},
			"MCD",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.output, moveCratesPart2(tc.moves, tc.stacks); want != got {
				t.Errorf("want: %v, got: %v", want, got)
			}
		})
	}
}
