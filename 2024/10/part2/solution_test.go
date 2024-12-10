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
				{".", ".", ".", ".", ".", "0", "."},
				{".", ".", "4", "3", "2", "1", "."},
				{".", ".", "5", ".", ".", "2", "."},
				{".", ".", "6", "5", "4", "3", "."},
				{".", ".", "7", ".", ".", "4", "."},
				{".", ".", "8", "7", "6", "5", "."},
				{".", ".", "9", ".", ".", ".", "."},
			},
			r: 0, c: 5,
			want: 3,
		},
		{
			name: "2",
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
			want: 13,
		},
		{
			name: "3",
			m: [][]string{
				{"0", "1", "2", "3", "4", "5"},
				{"1", "2", "3", "4", "5", "6"},
				{"2", "3", "4", "5", "6", "7"},
				{"3", "4", "5", "6", "7", "8"},
				{"4", ".", "6", "7", "8", "9"},
				{"5", "6", "7", "8", "9", "."},
			},
			r: 0, c: 0,
			want: 227,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ratings := map[pos]int{}
			getRatings(tc.m, tc.r, tc.c, ratings)
			var got int64
			for _, rating := range ratings {
				got += int64(rating)
			}

			if want := tc.want; want != got {
				t.Errorf("want=%v, got=%v", want, got)
			}
		})
	}

	t.Run("example_input_test", func(t *testing.T) {
		out, err := solution("../example_input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(81), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(1326), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})
}
