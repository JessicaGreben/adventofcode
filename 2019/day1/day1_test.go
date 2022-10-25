package day1_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jessicagreben/adventofcode/2019/day1"
	"github.com/jessicagreben/adventofcode/2019/pkg/input"
)

func TestCalculateFuelPart1(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	moduleMasses, err := input.ReadFromFile(filepath.Join(dir, "input.txt"))
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name  string
		input []string
		want  int
	}{
		{"1", []string{"1969"}, 654},
		{"2", []string{"100756"}, 33583},
		{"3", moduleMasses, 3317659},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fuel, err := day1.Part1(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if want, got := tc.want, fuel; want != got {
				t.Errorf("want: %d, got: %d", want, got)
			}
		})
	}
}

func TestCalculateFuelPart2(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	moduleMasses, err := input.ReadFromFile(filepath.Join(dir, "input.txt"))
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name  string
		input []string
		want  int
	}{
		{"1", []string{"1969"}, 966},
		{"2", []string{"100756"}, 50346},
		{"3", moduleMasses, 4973616},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fuel, err := day1.Part2(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if want, got := tc.want, fuel; want != got {
				t.Errorf("want: %d, got: %d", want, got)
			}
		})
	}
}
