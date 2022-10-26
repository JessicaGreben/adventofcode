package day1_test

import (
	"testing"

	"github.com/jessicagreben/adventofcode/2019/day1"
	"github.com/jessicagreben/adventofcode/2019/pkg/input"
)

func TestDay1Part1(t *testing.T) {
	moduleMasses, err := input.ReadFromFileToI("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{"1", []int{1969}, 654},
		{"2", []int{100756}, 33583},
		{"3", moduleMasses, 3317659},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, day1.Part1(tc.input); want != got {
				t.Errorf("want: %d, got: %d", want, got)
			}
		})
	}
}

func TestDay1Part2(t *testing.T) {
	moduleMasses, err := input.ReadFromFileToI("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{"1", []int{1969}, 966},
		{"2", []int{100756}, 50346},
		{"3", moduleMasses, 4973616},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, day1.Part2(tc.input); want != got {
				t.Errorf("want: %d, got: %d", want, got)
			}
		})
	}
}
