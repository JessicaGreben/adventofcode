package day6

import (
	"testing"
)

func TestPart1(t *testing.T) {
	if want, got := 520, part1(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestShortestPath(t *testing.T) {
	grid := parseInput("testinput.txt")
	if want, got := 31, shortestPath(grid); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
