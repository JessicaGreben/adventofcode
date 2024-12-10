package main

import "testing"

func TestSolutionPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int64
	}{
		{"1", "../input.txt", -1},
	}

	for _, tc := range testCases {
		t.Skip()
		t.Run(tc.name, func(t *testing.T) {
			out, err := solution(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if want, got := tc.want, out; want != got {
				t.Errorf("want=%v, got=%v", want, got)
			}
		})
	}

	t.Run("example_input_test", func(t *testing.T) {
		out, err := solution("../example_input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(-1), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		t.Skip()
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(-1), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})
}
