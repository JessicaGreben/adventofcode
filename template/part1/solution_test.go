package main

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"1", "", -1},
		{"2", "", -1},
		{"3", "", -1},
		{"4", "", -1},
		{"5", "", -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := processLine(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if want, got := tc.want, out; want != got {
				t.Errorf("want=%d, got=%d", want, got)
			}
		})
	}

	t.Run("input_test", func(t *testing.T) {
		out, err := solution("../input_test.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := 0, out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := 0, out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
}
