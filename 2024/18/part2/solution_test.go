package main

import (
	"fmt"
	"testing"
)

func TestSolutionPart2(t *testing.T) {
	t.Run("example_input_test", func(t *testing.T) {
		x, y, err := solution("../example_input.txt", 7, 12)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(x, y)
		if want, got := 6, y; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
		if want, got := 1, x; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		x, y, err := solution("../input.txt", 71, 1024)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(x, y)
		if want, got := 50, y; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
		if want, got := 28, x; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})
}
