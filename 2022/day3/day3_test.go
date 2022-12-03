package day1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	if want, got := 8493, part1(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	if want, got := 2552, part2(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestSumPriority2(t *testing.T) {

	testCases := []struct {
		name   string
		input  [][]string
		output int
	}{
		{
			name:   "1",
			input:  [][]string{{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}},
			output: 18,
		},
		{
			name:   "2",
			input:  [][]string{{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}},
			output: 52,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.output, sumPriority2(tc.input); want != got {
				t.Errorf("want: %d, got: %d", want, got)
			}
		})
	}
}
func TestSumPriority(t *testing.T) {
	testCases := []struct {
		name   string
		input  []string
		output int
	}{
		{"1", []string{"vJrwpWtwJgWrhcsFMMfFFhFp"}, 16},
		{"2", []string{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"}, 38},
		{"3", []string{"PmmdzqPrVvPwwTWBwg"}, 42},
		{"4", []string{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"}, 22},
		{"5", []string{"ttgJtRGJQctTZtZT"}, 20},
		{"6", []string{"CrZsJsPPZsGzwwsLwLmpwMDw"}, 19},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.output, sumPriority(tc.input); want != got {
				t.Errorf("want: %d, got: %d", want, got)
			}
		})
	}
}
