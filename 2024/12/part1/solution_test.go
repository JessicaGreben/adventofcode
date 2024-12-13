package main

import (
	"testing"

	"github.com/jessicagreben/adventofcode/pkg/matrix"
)

func TestSolutionPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input [][]string
		want  [][]int
	}{
		{
			"1",
			[][]string{
				{"0", "0", "0", "0", "0"},
				{"0", "X", "0", "X", "0"},
				{"0", "0", "0", "0", "0"},
				{"0", "X", "0", "X", "0"},
				{"0", "0", "0", "0", "0"},
			},
			[][]int{
				{1, 1, 1, 1, 1},
				{1, 2, 1, 3, 1},
				{1, 1, 1, 1, 1},
				{1, 4, 1, 5, 1},
				{1, 1, 1, 1, 1},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotMatrix := convert(tc.input)
			if len(gotMatrix) != len(tc.want) {
				t.Fatal("wrong size")
			}
			for r := range tc.want {
				for c := range tc.want[r] {
					if want, got := tc.want[r][c], gotMatrix[r][c]; want != got {
						matrix.Print(gotMatrix)
						t.Errorf("want=%v, got=%v", want, got)
					}
				}
			}
		})
	}

	t.Run("example_input_test", func(t *testing.T) {
		out, err := solution("../example_input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(1930), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(1437300), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})
}
