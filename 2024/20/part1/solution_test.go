package main

import "testing"

func TestSolutionPart1(t *testing.T) {
	t.Run("example_input_test", func(t *testing.T) {
		out, err := solution("../example_input.txt", -1)
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(44), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt", 100)
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(1375), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})
}
