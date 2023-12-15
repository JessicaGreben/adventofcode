package main

import (
	"testing"

	"github.com/jessicagreben/adventofcode/pkg/input"
)

func TestSolutionPart1(t *testing.T) {
	t.Run("ex_1_test", func(t *testing.T) {
		m, err := input.ConvertToMatrix("../ex1_input.txt")
		if err != nil {
			t.Fatal(err)
		}
		out := getVertical(m)
		if want, got := int64(5), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
		out = getHorizontal(m)
		if want, got := int64(-1), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
	t.Run("ex_2_test", func(t *testing.T) {
		m, err := input.ConvertToMatrix("../ex2_input.txt")
		if err != nil {
			t.Fatal(err)
		}
		out := getHorizontal(m)
		if want, got := int64(4), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
		out = getVertical(m)
		if want, got := int64(-1), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("ex_1_2_test", func(t *testing.T) {
		out, err := solution("../ex_1_2_input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(405), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(42974), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
}
