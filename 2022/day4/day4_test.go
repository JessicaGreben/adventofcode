package day1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	if want, got := 487, part1(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	if want, got := 849, part2(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestPartialOverlap(t *testing.T) {
	testCases := []struct {
		name   string
		input  pair
		output bool
	}{

		{
			"1",
			pair{a: []int{5, 7}, b: []int{7, 9}},
			true,
		},
		{
			"2",
			pair{a: []int{2, 8}, b: []int{3, 7}},
			true,
		},
		{
			"3",
			pair{a: []int{6, 6}, b: []int{4, 6}},
			true,
		},
		{
			"4",
			pair{a: []int{2, 6}, b: []int{4, 8}},
			true,
		},
		{
			"5",
			pair{a: []int{2, 6}, b: []int{1, 3}},
			true,
		},
		{
			"6",
			pair{a: []int{2, 6}, b: []int{7, 8}},
			false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.output, tc.input.partialoverlap(); want != got {
				t.Errorf("want: %v, got: %v", want, got)
			}
		})
	}
}
func TestFullOverlap(t *testing.T) {
	testCases := []struct {
		name   string
		input  pair
		output bool
	}{

		{
			"1",
			pair{a: []int{2, 8}, b: []int{3, 7}},
			true,
		},
		{
			"2",
			pair{a: []int{6, 6}, b: []int{4, 6}},
			true,
		},
		{
			"3",
			pair{a: []int{7, 8}, b: []int{4, 6}},
			false,
		},
		{
			"4",
			pair{a: []int{1, 2}, b: []int{4, 6}},
			false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.output, tc.input.fulloverlap(); want != got {
				t.Errorf("want: %v, got: %v", want, got)
			}
		})
	}
}
