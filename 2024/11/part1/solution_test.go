package main

import "testing"

func TestSolutionPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input []int64
		want  []int64
	}{
		{"1", []int64{125, 17}, []int64{253000, 1, 7}},
		{"2", []int64{253000, 1, 7}, []int64{253, 0, 2024, 14168}},
		{"3", []int64{253, 0, 2024, 14168}, []int64{512072, 1, 20, 24, 28676032}},
		{"4", []int64{512072, 1, 20, 24, 28676032}, []int64{512, 72, 2024, 2, 0, 2, 4, 2867, 6032}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := blink(tc.input)
			if want, got := len(tc.want), len(got); want != got {
				t.Fatalf("len want=%v, got=%v", want, got)
			}
			for i := range tc.want {
				if want, got := tc.want[i], got[i]; want != got {
					t.Errorf("want=%v, got=%v", want, got)
				}
			}
		})
	}

	t.Run("example_input_test", func(t *testing.T) {
		out, err := solution("../example_input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(55312), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(209412), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})
}
