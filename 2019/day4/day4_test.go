package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestPart1AoC(t *testing.T) {
	if want, got := 1955, day4part1(); got != want {
		t.Errorf("AoC part 1: want: %d, got: %d", want, got)
	}
}

func TestPart2AoC(t *testing.T) {
	if want, got := 1319, day4part2(); got != want {
		t.Errorf("AoC part 1: want: %d, got: %d", want, got)
	}
}
func TestIsValidPasswordPart1(t *testing.T) {
	input := 666640
	intDigits := strings.Split(strconv.Itoa(input), "")
	if want, got := []string{"6", "6", "6", "6", "4", "0"}, intDigits; len(want) != len(got) {
		t.Fatalf("got: %v, want: %v", got, want)
	}
	if want, got := false, isAsc(intDigits); want != got {
		t.Errorf("got: %v, want: %v", got, want)
	}
	if want, got := true, twoAdjNums(intDigits); want != got {
		t.Errorf("got: %v, want: %v", got, want)
	}

	if want, got := true, isValidPasswordPart1(111111); want != got {
		t.Fatalf("got: %v, want: %v", got, want)
	}
	if want, got := false, isValidPasswordPart1(223450); want != got {
		t.Fatalf("got: %v, want: %v", got, want)
	}
	if want, got := false, isValidPasswordPart1(123789); want != got {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}
func TestIsValidPasswordPart2(t *testing.T) {
	input := 666640
	intDigits := strings.Split(strconv.Itoa(input), "")
	if want, got := []string{"6", "6", "6", "6", "4", "0"}, intDigits; len(want) != len(got) {
		t.Fatalf("got: %v, want: %v", got, want)
	}
	if want, got := false, isAsc(intDigits); want != got {
		t.Errorf("got: %v, want: %v", got, want)
	}
	if want, got := true, twoAdjNums(intDigits); want != got {
		t.Errorf("got: %v, want: %v", got, want)
	}

	if want, got := map[string]bool{"6": true}, occurTwoAdjNums(intDigits); len(want) != len(got) {
		t.Fatalf("got: %v, want: %v", got, want)
	}
	if want, got := map[string]bool{"6": true}, occurThreeAdjNums(intDigits); len(want) != len(got) {
		t.Fatalf("got: %v, want: %v", got, want)
	}
	if want, got := false, isValidPasswordPart2(666640); want != got {
		t.Fatalf("got: %v, want: %v", got, want)
	}
	if want, got := false, isValidPasswordPart2(126666); want != got {
		t.Fatalf("got: %v, want: %v", got, want)
	}
	if want, got := true, isValidPasswordPart2(123466); want != got {
		t.Fatalf("got: %v, want: %v", got, want)
	}
	if want, got := true, isValidPasswordPart2(111122); want != got {
		t.Fatalf("got: %v, want: %v", got, want)
	}
	if want, got := false, isValidPasswordPart2(111222); want != got {
		t.Fatalf("got: %v, want: %v", got, want)
	}
	if want, got := true, isValidPasswordPart2(112233); want != got {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}
