package day5

import (
	"testing"
)

func TestProgramPart1(t *testing.T) {
	output := TESTDiagnosticProgram(1)
	want := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 14522484}

	if want, got := len(want), len(output); want != got {
		t.Fatalf("len want: %d, got: %d", want, got)
	}

	for i, out := range output {
		if want, got := want[i], out; want != got {
			t.Errorf("item %d want: %d, got: %d", i, want, got)
		}
	}
}
