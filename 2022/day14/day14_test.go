package day14

import "testing"

func TestPart1(t *testing.T) {
	if want, got := 592, part1(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestPart1Example(t *testing.T) {
	pathPoints, minY := convertLines(readLines("testinput.txt"))
	if want, got := 24, countSand(pathPoints, minY); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	if want, got := 30367, part2(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestPart2Example(t *testing.T) {
	pathPoints, minY := convertLines(readLines("testinput.txt"))
	if want, got := 93, countSandPart2(pathPoints, minY); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
