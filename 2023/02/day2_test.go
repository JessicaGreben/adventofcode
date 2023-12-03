package main

import "testing"

func TestDay02Part1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"1", "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 1},
		{"2", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 2},
		{"3", "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 0},
		{"4", "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 0},
		{"5", "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := doPart1(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if want, got := tc.want, got; want != got {
				t.Errorf("want=%d, got=%d", want, got)
			}
		})
	}

	sum, err := readInput("input_test.txt", doPart1)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := 8, sum; want != got {
		t.Errorf("total sum want=%d, got=%d", want, got)
	}

	sum, err = readInput("", doPart1)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := 2149, sum; want != got {
		t.Errorf("total sum want=%d, got=%d", want, got)
	}
}

func TestDay02Part2(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"1", "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 48},
		{"2", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12},
		{"3", "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 1560},
		{"4", "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 630},
		{"5", "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 36},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := doPart2(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if want, got := tc.want, got; want != got {
				t.Errorf("want=%d, got=%d", want, got)
			}
		})
	}

	sum, err := readInput("input_test.txt", doPart2)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := 2286, sum; want != got {
		t.Errorf("total sum want=%d, got=%d", want, got)
	}

	sum, err = readInput("", doPart2)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := 71274, sum; want != got {
		t.Errorf("total sum want=%d, got=%d", want, got)
	}
}
