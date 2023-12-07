package main

import "testing"

func TestSolutionPart2(t *testing.T) {
	t.Run("input_test", func(t *testing.T) {
		races := []race{
			{71530, 940200},
		}
		if want, got := int64(71503), solution(races); want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		races := []race{
			{46689866, 358105418071080},
		}

		if want, got := int64(-1), solution(races); want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
}
