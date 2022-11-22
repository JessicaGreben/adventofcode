package day2_test

import (
	"testing"

	"golang.org/x/exp/slices"

	"github.com/jessicagreben/adventofcode/2019/day2"
	"github.com/jessicagreben/adventofcode/2019/pkg/intcode"
)

func TestDay2Intcode(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output []int
		want   int
	}{
		{"1", []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}, 2},
		{"2", []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}, 2},
		{"3", []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}, 2},
		{"4", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, 30},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := intcode.NewProgram(tc.input)
			solution, err := p.Run(-1)
			if err != nil {
				t.Fatal(err)
			}
			if want, got := tc.want, solution; want != got {
				t.Errorf("want: %d, got: %d", want, got)
			}

			for i := range tc.output {
				if want, got := tc.output[i], tc.input[i]; want != got {
					t.Errorf("want: %d, got: %d", want, got)
				}
			}
		})
	}
}

func TestDay2Part1(t *testing.T) {
	solution, err := day2.Part1(slices.Clone(day2.AoCInput))
	if err != nil {
		t.Fatal(err)
	}
	if want, got := 2842648, solution; want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestDay2Part2(t *testing.T) {
	noun, verb, err := day2.Part2(slices.Clone(day2.AoCInput))
	if err != nil {
		t.Fatal(err)
	}
	if want, got := 9074, 100*noun+verb; want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
