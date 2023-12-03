package input_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jessicagreben/adventofcode/pkg/input"
)

func TestConvertToMatrix(t *testing.T) {
	m, err := input.ConvertToMatrix("matrix_test_input.txt")
	if err != nil {
		t.Fatal(err)
	}
	if want, got := []string{"4", "6", ".", ".", "1", "4", "."}, m[0]; !cmp.Equal(want, got) {
		t.Errorf("(-want +got):\n%s", cmp.Diff(want, got))
	}
	if want, got := []string{".", ".", ".", "*", ".", ".", "."}, m[1]; !cmp.Equal(want, got) {
		t.Errorf("(-want +got):\n%s", cmp.Diff(want, got))
	}
	if want, got := []string{".", ".", "3", "5", ".", "6", "3"}, m[2]; !cmp.Equal(want, got) {
		t.Errorf("(-want +got):\n%s", cmp.Diff(want, got))
	}
	if want, got := []string{".", ".", ".", ".", "#", ".", "."}, m[3]; !cmp.Equal(want, got) {
		t.Errorf("(-want +got):\n%s", cmp.Diff(want, got))
	}
}
