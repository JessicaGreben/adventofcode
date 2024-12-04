package main

import "testing"

func TestSolutionPart1(t *testing.T) {
	t.Run("input_test", func(t *testing.T) {
		out, err := solution("../example_input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(18), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(2532), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
}
