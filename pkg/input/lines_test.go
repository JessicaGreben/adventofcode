package input_test

import (
	"strconv"
	"testing"

	"github.com/jessicagreben/adventofcode/pkg/input"
)

func TestForEachLine(t *testing.T) {
	out, err := input.ForEachLine("lines_test_input.txt", func(line string) (int, error) {
		num, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		return num, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if want, got := out, 15; want != got {
		t.Errorf("want=%d, got=%d", want, got)
	}
}
