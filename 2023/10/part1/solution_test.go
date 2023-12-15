package main

import (
	"testing"
)

func TestSolutionPart1(t *testing.T) {
	t.Run("simple_example", func(t *testing.T) {
		grid := [][]string{
			{".", ".", ".", ".", "."},
			{".", "S", "-", "7", "."},
			{".", "|", ".", "|", "."},
			{".", "L", "-", "J", "."},
			{".", ".", ".", ".", "."},
		}
		g := generateGraph(grid)
		g.s.pipeType = "F"
		g.traverse(g.s)
		if want, got := int64(4), g.maxDistance/2+1; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("simple_input_test", func(t *testing.T) {
		out, err := solution("../simple_input_test.txt", "F")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(4), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
	t.Run("complex_input_test", func(t *testing.T) {
		out, err := solution("../complex_input_test.txt", "F")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(8), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt", "L")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(6886), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
}
