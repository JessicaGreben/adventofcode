package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolutionPart2(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"1", "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 4},
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

	t.Run("foreachline", func(t *testing.T) {
		cardToWinCount, err := ForEachLine("../input_test.txt", processLine)
		if err != nil {
			t.Fatal(err)
		}
		if want, got := 6, len(cardToWinCount); want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
		expected := map[int]int{
			1: 4, 2: 2, 3: 2, 4: 1, 5: 0, 6: 0,
		}
		if want, got := expected, cardToWinCount; !cmp.Equal(want, got) {
			t.Errorf("(-want +got):\n%s", cmp.Diff(want, got))
		}
	})

	t.Run("total card count", func(t *testing.T) {
		winCards := map[int]int{
			1: 4, 2: 2, 3: 2, 4: 1, 5: 0, 6: 0,
		}
		cardCount, err := totalCardCount(winCards)
		if err != nil {
			t.Fatal(err)
		}
		if want, got := 6, len(cardCount); want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
		expected := map[int]int{
			1: 1, 2: 2, 3: 4, 4: 8, 5: 14, 6: 1,
		}
		if want, got := expected, cardCount; !cmp.Equal(want, got) {
			t.Errorf("(-want +got):\n%s", cmp.Diff(want, got))
		}
	})

	t.Run("input_test", func(t *testing.T) {
		out, err := solution("../input_test.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := 30, out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := 5659035, out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
}
