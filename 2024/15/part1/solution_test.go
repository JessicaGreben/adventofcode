package main

import (
	"testing"
)

func TestSolutionPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input [][]string
		moves string
		want  [][]string
	}{
		{
			name: "1",
			input: [][]string{
				{"#", "#", "#", "#", "#"},
				{"#", "@", "O", ".", "#"},
				{"#", "#", "#", "#", "#"},
			},
			moves: ">",
			want: [][]string{
				{"#", "#", "#", "#", "#"},
				{"#", ".", "@", "O", "#"},
				{"#", "#", "#", "#", "#"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			traverse(tc.input, tc.moves, 1, 1, 0)
			for r := range tc.want {
				for c := range tc.want[r] {
					if want, got := tc.want[r][c], tc.input[r][c]; want != got {
						t.Errorf("r=%v, c=%v, want=%v, got=%v", r, c, want, got)
					}
				}
			}
		})
	}

	t.Run("small_example_input_test", func(t *testing.T) {
		moves := "<^^>>>vv<v>>v<<"
		out, err := solution("../small_example_input.txt", moves)
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(2028), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})

	t.Run("example_input_test", func(t *testing.T) {
		out, err := solution("../example_input.txt", exampleMoves)
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(10092), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt", inputMoves)
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(1517819), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})
}
