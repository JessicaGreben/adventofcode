package day1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	// too low
	if want, got := 5675922, part1(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestProgramPart1(t *testing.T) {
	in, err := input("testinput.txt")
	if err != nil {
		t.Fatal(err)
	}

	if want, got := 26, countNoBeacons(in, 10); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestCount(t *testing.T) {
	testCase := []struct {
		name string
		y    int
		out  int
	}{
		{"1", 7 + 0, 19},
		{"2", 7 + 3, 12},
		{"3", 7 + 9, 1},
		{"4", 7 - 9, 1},
		{"5", 7 + 10, 0},
		{"6", 7 - 8, 3},
	}

	in := []sensor{{point{8, 7}, point{2, 10}}}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.out, countNoBeacons(in, tc.y); want != got {
				t.Errorf("want: %d, got: %d", want, got)
			}
		})
	}
}
