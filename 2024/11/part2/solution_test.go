package main

import "testing"

func TestSolutionPart2(t *testing.T) {
	testCases := []struct {
		name  string
		input int64
		want  int64
	}{
		{"1", 125, 7},
		{"2", 17, 15},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			memo := map[memoKey]int64{}
			if want, got := tc.want, dfs(tc.input, 0, 6, 1, memo); want != got {
				t.Errorf("want=%v, got=%v", want, got)
			}
		})
	}

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(248967696501656), out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})
}
