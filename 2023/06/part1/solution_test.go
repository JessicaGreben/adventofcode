package main

import "testing"

func TestSolutionPart1(t *testing.T) {
	t.Run("input_test", func(t *testing.T) {
		races := []race{
			{7, 9},
			{15, 40},
			{30, 200},
		}
		if want, got := 288, solution(races); want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		races := []race{
			{46, 358},
			{68, 1054},
			{98, 1807},
			{66, 1080},
		}

		if want, got := 138915, solution(races); want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}

	})
}
