package main

import "testing"

func TestSolutionPart2(t *testing.T) {
	testCases := []struct {
		name string
		m    [][]string
		r, c int
		dir  pos
		want bool
	}{
		{
			name: "1",
			m: [][]string{
				{"M", ".", "S"},
				{".", "A", "."},
				{"M", ".", "S"},
			},
			r: 0, c: 0,
			dir:  pos{1, 1},
			want: true,
		},
		{
			name: "2",
			m: [][]string{
				{"M", ".", "M"},
				{".", "A", "."},
				{"S", ".", "S"},
			},
			r: 0, c: 0,
			dir:  pos{1, 1},
			want: true,
		},
		{
			name: "3",
			m: [][]string{
				{"M", ".", "M"},
				{".", "A", "."},
				{"S", ".", "S"},
			},
			r: 0, c: 2,
			dir:  pos{1, -1},
			want: true,
		},
		{
			name: "4",
			m: [][]string{
				{"S", ".", "M"},
				{".", "A", "."},
				{"S", ".", "M"},
			},
			r: 0, c: 2,
			dir:  pos{1, -1},
			want: true,
		},
		{
			name: "5",
			m: [][]string{
				{"M", ".", "S"},
				{".", "A", "."},
				{"M", ".", "S"},
			},
			r: 2, c: 0,
			dir:  pos{-1, 1},
			want: true,
		},
		{
			name: "6",
			m: [][]string{
				{"S", ".", "S"},
				{".", "A", "."},
				{"M", ".", "M"},
			},
			r: 2, c: 0,
			dir:  pos{-1, 1},
			want: true,
		},
		{
			name: "7",
			m: [][]string{
				{"S", ".", "S"},
				{".", "A", "."},
				{"M", ".", "M"},
			},
			r: 2, c: 2,
			dir:  pos{-1, -1},
			want: true,
		},
		{
			name: "8",
			m: [][]string{
				{"S", ".", "M"},
				{".", "A", "."},
				{"S", ".", "M"},
			},
			r: 2, c: 2,
			dir:  pos{-1, -1},
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, isXmas(tc.m, tc.r, tc.c, tc.dir); want != got {
				t.Errorf("want=%v, got=%v", want, got)
			}
		})
	}

	t.Run("input_test", func(t *testing.T) {
		out, err := solution("../example_input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(9), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(1941), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
}
