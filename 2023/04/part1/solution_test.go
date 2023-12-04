package main

import "testing"

func TestSolutionPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"1", "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 8},
		{"2", "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 2},
		{"3", "Card 3: 1 21 53 59 44 | 69 82 63 72 16 21 14  1", 2},
		{"4", "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", 1},
		{"5", "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", 0},
		{"6", "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 0},
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
		if want, got := 13, out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := 24160, out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
}
