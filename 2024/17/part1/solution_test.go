package main

import "testing"

func TestSolutionPart1(t *testing.T) {
	t.Run("example_input_test", func(t *testing.T) {
		t.Skip()
		/*
			Register A: 729
			Register B: 0
			Register C: 0
			Program: 0,1,5,4,3,0
		*/
		out, err := solution(729, 0, 0, []int{0, 1, 5, 4, 3, 0})
		if err != nil {
			t.Fatal(err)
		}
		if want, got := "4,6,3,5,6,3,5,2,1,0", out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		/*
			Register A: 52884621
			Register B: 0
			Register C: 0
			Program: 2,4,1,3,7,5,4,7,0,3,1,5,5,5,3,0
		*/
		out, err := solution(52884621, 0, 0, []int{2, 4, 1, 3, 7, 5, 4, 7, 0, 3, 1, 5, 5, 5, 3, 0})
		if err != nil {
			t.Fatal(err)
		}
		if want, got := "1,3,5,1,7,2,5,1,6", out; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	})
}
