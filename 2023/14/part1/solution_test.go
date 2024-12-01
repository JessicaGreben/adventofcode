package main

import (
	"testing"

	"github.com/jessicagreben/adventofcode/pkg/input"
)

func TestSolutionPart1(t *testing.T) {
	t.Run("input_test", func(t *testing.T) {
		out, err := solution("../input_test.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(136), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		t.Skip()
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(0), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
}

func TestSolutionTilt(t *testing.T) {
	m, err := input.ConvertToMatrix("../input_test.txt")
	if err != nil {
		t.Fatal(err)
	}

	tiltNorth(m)
	want := [][]string{
		{"O", "O", "O", "O", ".", "#", ".", "O", ".", "."},
		{"O", "O", ".", ".", "#", ".", ".", ".", ".", "#"},
		{"O", "O", ".", ".", "O", "#", "#", ".", ".", "O"},
		{"O", ".", ".", "#", ".", "O", "O", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
		{".", ".", "#", ".", ".", ".", ".", "#", ".", "#"},
		{".", ".", "O", ".", ".", "#", ".", "O", ".", "O"},
		{".", ".", "O", ".", ".", ".", ".", ".", ".", "."},
		{"#", ".", ".", ".", ".", "#", "#", "#", ".", "."},
		{"#", ".", ".", ".", ".", "#", ".", ".", ".", "."},
	}
	if want, got := len(want), len(m); want != got {
		t.Errorf("want=%d, got=%d", want, got)
	}

	for i := range want {
		wantRow, gotRow := want[i], m[i]
		for j := range wantRow {
			if want, got := wantRow[j], gotRow[j]; want != got {
				t.Errorf("want=%s, got=%s", want, got)
			}
		}
	}
}
