package day1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	if want, got := 8890, part1(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	if want, got := 10238, part2(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestScore(t *testing.T) {
	input := []string{"A Y", "B X", "C Z"}
	if want, got := 15, totalScore(input); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
