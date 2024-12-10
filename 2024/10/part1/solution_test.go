package main

import "testing"

func TestSolutionPart1(t *testing.T) {
	testCases := []struct {
		name string
		m    [][]string
		r, c int
		want int64
	}{
		{
			name: "1",
			m: [][]string{
				{"0", "1", "2", "3"},
				{"1", "2", "3", "4"},
				{"8", "7", "6", "5"},
				{"9", "8", "7", "6"},
			},
			r: 0, c: 0,
			want: 1,
		},
		{
			name: "2",
			m: [][]string{
				{".", ".", ".", "0", ".", ".", "."},
				{".", ".", ".", "1", ".", ".", "."},
				{".", ".", ".", "2", ".", ".", "."},
				{"6", "5", "4", "3", "4", "5", "6"},
				{"7", ".", ".", ".", ".", ".", "7"},
				{"8", ".", ".", ".", ".", ".", "8"},
				{"9", ".", ".", ".", ".", ".", "9"},
			},
			r: 0, c: 3,
			want: 2,
		},
		{
			name: "3",
			m: [][]string{
				{".", ".", "9", "0", ".", ".", "9"},
				{".", ".", ".", "1", ".", "9", "8"},
				{".", ".", ".", "2", ".", ".", "7"},
				{"6", "5", "4", "3", "4", "5", "6"},
				{"7", "6", "5", ".", "9", "8", "7"},
				{"8", "7", "6", ".", ".", ".", "."},
				{"9", "8", "7", ".", ".", ".", "."},
			},
			r: 0, c: 3,
			want: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			scores := map[pos]bool{}
			getScore(tc.m, tc.r, tc.c, scores)
			if want, got := tc.want, int64(len(scores)); want != got {
				t.Errorf("want=%v, got=%v", want, got)
			}
		})
	}

	t.Run("example_input_test", func(t *testing.T) {
		out, err := solution("../example_input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(36), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(709), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})
}
