package day1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	cals, err := mostCalories()
	if err != nil {
		t.Fatal(err)
	}
	if want, got := 71471, cals; want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	cals, err := top3Calories()
	if err != nil {
		t.Fatal(err)
	}
	if want, got := 211189, cals; want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestInsert(t *testing.T) {
	testCases := []struct {
		name   string
		arr    []int
		value  int
		output []int
	}{
		{"1", []int{0, 0, 0}, 1, []int{1, 0, 0}},
		{"2", []int{1, 0, 0}, 1, []int{1, 1, 0}},
		{"3", []int{1, 1, 0}, 2, []int{1, 1, 2}},
		{"4", []int{2, 5, 9}, 1, []int{2, 5, 9}},
		{"5", []int{2, 5, 9}, 10, []int{10, 5, 9}},
		{"6", []int{2, 5, 9}, 3, []int{3, 5, 9}},
		{"7", []int{2, 5, 9}, 7, []int{7, 5, 9}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := insert(tc.arr, tc.value)
			if want, got := tc.output, out; len(want) != len(out) {
				t.Fatalf("len want=%v, got=%v", want, got)
			}

			for i := range out {
				if want, got := tc.output, out; want[i] != got[i] {
					t.Errorf("want=%d, got=%d", want, got)
				}
			}
		})
	}
}
