package day8

import (
	"testing"
)

func TestPart1(t *testing.T) {
	if want, got := 1859, part1(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	if want, got := 332640, part2(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestVisible(t *testing.T) {
	grid := [][]tree{
		{{height: 3}, {height: 0}, {height: 3}, {height: 7}, {height: 3}},
		{{height: 2}, {height: 5}, {height: 5}, {height: 1}, {height: 2}},
		{{height: 6}, {height: 5}, {height: 3}, {height: 3}, {height: 2}},
		{{height: 3}, {height: 3}, {height: 5}, {height: 4}, {height: 9}},
		{{height: 3}, {height: 5}, {height: 3}, {height: 9}, {height: 0}},
	}
	if want, got := 21, findVisible(grid); want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestRowVisible(t *testing.T) {
	grid := [][]tree{
		{{height: 3}, {height: 0}, {height: 3}, {height: 7}, {height: 3}},
		{{height: 2}, {height: 5}, {height: 5}, {height: 1}, {height: 2}},
		{{height: 6}, {height: 5}, {height: 3}, {height: 3}, {height: 2}},
		{{height: 3}, {height: 3}, {height: 5}, {height: 4}, {height: 9}},
		{{height: 3}, {height: 5}, {height: 3}, {height: 9}, {height: 0}},
	}

	testCases := []struct {
		name                   string
		r, c                   int
		edge                   bool
		rowShorter, colShorter bool
	}{
		{
			"1",
			1, 1,
			false, true, true,
		},
		{
			"2",
			1, 2,
			false, true, true,
		},
		{
			"3",
			1, 3,
			false, false, false,
		},
		{
			"4",
			2, 1,
			false, true, false,
		},
		{
			"5",
			2, 2,
			false, false, false,
		},
		{
			"6",
			2, 3,
			false, true, false,
		},
		{
			"7",
			3, 1,
			false, false, false,
		},
		{
			"8",
			3, 2,
			false, true, true,
		},
		{
			"9",
			3, 3,
			false, false, false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.edge, isEdge(tc.r, tc.c, grid); want != got {
				t.Errorf("want: %v, got: %v", want, got)
			}
			if want, got := tc.rowShorter, rowShorter(tc.r, tc.c, grid); want != got {
				t.Errorf("want: %v, got: %v", want, got)
			}
			if want, got := tc.colShorter, colShorter(tc.r, tc.c, grid); want != got {
				t.Errorf("want: %v, got: %v", want, got)
			}
		})
	}
}

func TestScore(t *testing.T) {
	grid := [][]tree{
		{{height: 3}, {height: 0}, {height: 3}, {height: 7}, {height: 3}},
		{{height: 2}, {height: 5}, {height: 5}, {height: 1}, {height: 2}},
		{{height: 6}, {height: 5}, {height: 3}, {height: 3}, {height: 2}},
		{{height: 3}, {height: 3}, {height: 5}, {height: 4}, {height: 9}},
		{{height: 3}, {height: 5}, {height: 3}, {height: 9}, {height: 0}},
	}
	if want, got := 8, bestScenicScore(grid); want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestRowScore(t *testing.T) {
	grid := [][]tree{
		{{height: 3}, {height: 0}, {height: 3}, {height: 7}, {height: 3}},
		{{height: 2}, {height: 5}, {height: 5}, {height: 1}, {height: 2}},
		{{height: 6}, {height: 5}, {height: 3}, {height: 3}, {height: 2}},
		{{height: 3}, {height: 3}, {height: 5}, {height: 4}, {height: 9}},
		{{height: 3}, {height: 5}, {height: 3}, {height: 9}, {height: 0}},
	}

	testCases := []struct {
		name               string
		r, c               int
		rowScore, colScore int
	}{
		{
			"1",
			1, 2,
			1 * 2, 1 * 2,
		},
		{
			"2",
			3, 2,
			2 * 2, 2 * 1,
		},
		{
			"3",
			0, 0,
			0, 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.rowScore, rowScore(tc.r, tc.c, grid); want != got {
				t.Errorf("want: %v, got: %v", want, got)
			}
			if want, got := tc.colScore, colScore(tc.r, tc.c, grid); want != got {
				t.Errorf("want: %v, got: %v", want, got)
			}
		})
	}
}
