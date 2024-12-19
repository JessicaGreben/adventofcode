package main

import (
	"testing"
)

func TestSolutionPart2(t *testing.T) {
	t.Run("example_input_test", func(t *testing.T) {
		out, err := solution("../example_input.txt", exampleInputTowels)
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(16), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt", inputTowels)
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(636483903099279), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})
}
