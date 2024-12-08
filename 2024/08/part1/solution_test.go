package main

import "testing"

func TestSolutionPart1(t *testing.T) {
	testCases := []struct {
		name string
		m    [][]string
		a, b pos
		want []pos
	}{
		{
			name: "1",
			m: [][]string{
				{".", "x", "x", "."},
			},
			a:    pos{0, 1},
			b:    pos{0, 2},
			want: []pos{{0, 0}, {0, 3}},
		},
		{
			name: "2",
			m: [][]string{
				{".", ".", ""},
				{".", "x", ""},
				{".", "x", ""},
				{".", ".", ""},
			},
			a:    pos{1, 1},
			b:    pos{2, 1},
			want: []pos{{0, 1}, {3, 1}},
		},
		{
			name: "3",
			m: [][]string{
				{".", ".", "", ""},
				{".", "x", "", ""},
				{".", ".", "x", ""},
				{".", ".", "", ""},
			},
			a:    pos{1, 1},
			b:    pos{2, 2},
			want: []pos{{0, 0}, {3, 3}},
		},
		{
			name: "4",
			m: [][]string{
				{".", ".", "", ""},
				{".", ".", "x", ""},
				{".", "x", "", ""},
				{".", ".", "", ""},
			},
			a:    pos{1, 2},
			b:    pos{2, 1},
			want: []pos{{0, 3}, {3, 0}},
		},
		{
			name: "5",
			m: [][]string{
				{".", ".", "", ""},
				{".", ".", "", "x"},
				{".", "x", "", ""},
				{".", ".", "", ""},
			},
			a:    pos{1, 3},
			b:    pos{2, 1},
			want: []pos{},
		},
		{
			name: "6",
			m: [][]string{
				{"x", ".", "", "", "", ""},
				{".", ".", "", "", "", ""},
				{".", ".", "x", "", "", ""},
				{".", ".", "", "", "", ""},
				{".", ".", "", "", "", ""},
			},
			a:    pos{0, 0},
			b:    pos{2, 2},
			want: []pos{{4, 4}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := getAntinodes(tc.m, tc.a, tc.b)
			if want, got := len(tc.want), len(got); want != got {
				t.Fatalf("wrong len want=%v, got=%v", want, got)
			}
			for i := range tc.want {
				if want, got := tc.want[i], got[i]; want != got {
					t.Errorf("i=%v want=%v, got=%v", i, want, got)
				}
			}
		})
	}

	t.Run("input_test", func(t *testing.T) {
		out, err := solution("../example_input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(14), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(259), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})
}
