package main

import (
	"testing"
)

func TestSolutionPart2(t *testing.T) {
	t.Run("input_test", func(t *testing.T) {
		out, err := solution("../input_test.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(46), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		//t.Skip()
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(99751240), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
}
