package day6

import (
	"testing"
)

func TestPart1(t *testing.T) {
	if want, got := 1343, part1(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestPacketMarker(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output int
	}{

		{
			"1",
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			7,
		},
		{
			"2",
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			5,
		},
		{
			"3",
			"nppdvjthqldpwncqszvftbrmjlhg",
			6,
		},
		{
			"4",
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			10,
		},
		{
			"5",
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			11,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.output, findPacketMarker(tc.input); want != got {
				t.Errorf("want: %v, got: %v", want, got)
			}
		})
	}
}
