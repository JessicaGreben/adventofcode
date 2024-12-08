package main

import "testing"

func TestSolutionPart2(t *testing.T) {
	t.Run("input_test", func(t *testing.T) {
		out, err := solution("../example_input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(34), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(927), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})
}
