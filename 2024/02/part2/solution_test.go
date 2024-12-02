package main

import "testing"

func TestSolutionPart2(t *testing.T) {
	testCases := []struct {
		name    string
		input   []int64
		skipIdx int
		want    bool
	}{
		{"1", []int64{1, 2}, 1, true},
		{"2", []int64{1, 5}, 0, true},
		{"3", []int64{1, 1}, 1, true},
		{"4", []int64{1, 2, 3}, -1, true},
		{"5", []int64{1, 5, 2}, 1, true},
		{"6", []int64{1, 5, 1}, 1, false},
		{"7", []int64{1, 1, 5}, -1, false},
		{"8", []int64{1, 1, 2, -1}, -1, false},
		{"9", []int64{3, 1, 2, 3}, 0, true},
		{"10", []int64{3, 1, 2, 3, 4}, 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, isSafe(tc.input, tc.skipIdx); want != got {
				t.Errorf("want=%v, got=%v", want, got)
			}
		})
	}
	t.Run("input_test", func(t *testing.T) {
		out, err := solution("../example_input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(4), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(271), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
}
