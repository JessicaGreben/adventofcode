package main

import "testing"

func TestSolutionPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int64
	}{
		{"1", "0 3 6 9 12 15", 18},
		{"2", "1 3 6 10 15 21", 28},
		{"3", "10 13 16 21 30 45", 68},
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
		if want, got := int64(114), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(2005352194), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
}
