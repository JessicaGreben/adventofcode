package day17

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	if want, got := 3124, part1(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestPart1Example(t *testing.T) {
	lines := readLines("testinput.txt")
	moves := strings.Split(lines[0], "")
	if want, got := 3068, fallingRocks(2022, moves); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
