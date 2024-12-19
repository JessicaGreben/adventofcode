package main

import (
	"strings"
	"testing"

	"github.com/jessicagreben/adventofcode/pkg/trie"
)

func TestSolutionPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{"1", "brwrr", true},
		{"2", "bggr", true},
		{"3", "gbbr", true},
		{"4", "rrbgbr", true},
		{"5", "ubwu", false},
		{"6", "bwurrg", true},
		{"7", "brgr", true},
		{"8", "bbrgwb", false},
	}
	prefixTree := trie.New()
	for _, x := range strings.Split(exampleInputTowels, ", ") {
		prefixTree.Insert(x)
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, isDesignPossible(tc.input, 0, len(tc.input), prefixTree); want != got {
				t.Errorf("want=%v, got=%v", want, got)
			}
		})
	}

	t.Run("example_input_test", func(t *testing.T) {
		out, err := solution("../example_input.txt", exampleInputTowels)
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(6), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt", inputTowels)
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(285), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})
}
