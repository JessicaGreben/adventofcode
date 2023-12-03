package main

import "testing"

func TestDay02Part1(t *testing.T) {
	m, err := convertInputToMatrix("input_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(m); i++ {
		if want, got := len(m[0]), len(m[i]); want != got {
			t.Errorf("input matrix rows are not the same len, want=%d, got=%d", want, got)
		}
	}
	solution, err := doPart1(m)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := 4361, solution; want != got {
		t.Errorf("want=%d, got=%d", want, got)
	}

	m, err = convertInputToMatrix("")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(m); i++ {
		if want, got := len(m[0]), len(m[i]); want != got {
			t.Errorf("input matrix rows are not the same len, want=%d, got=%d", want, got)
		}
	}
	solution, err = doPart1(m)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := 532428, solution; want != got {
		t.Errorf("want=%d, got=%d", want, got)
	}
}

func TestGetNum(t *testing.T) {
	m, err := convertInputToMatrix("input_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	testCases := []struct {
		name string
		r    int
		c    int
		want int
	}{
		{"first num", 0, 0, 467},
		{"second num", 0, 1, -1},
		{"third num", 0, 2, -1},
	}

	seen = map[pos]bool{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := getNum(m, tc.r, tc.c)
			if err != nil {
				t.Fatal(err)
			}
			if want, got := tc.want, got; want != got {
				t.Errorf("want=%d, got=%d", want, got)
			}
		})
	}
}

func TestgetNeighboringNums(t *testing.T) {
	m, err := convertInputToMatrix("input_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	testCases := []struct {
		name     string
		r        int
		c        int
		wantNum  int
		wantNums []int
	}{
		{"first * symbol", 1, 3, 502, []int{467, 35}},
		{"second * symbol", 4, 3, 617, []int{617}},
		{"first # symbol", 3, 6, 633, []int{633}},
		{"+ symbol", 5, 5, 592, []int{592}},
		{"$ symbol", 8, 3, 664, []int{664}},
	}

	seen = map[pos]bool{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sum, nums, err := getNeighboringNums(m, tc.r, tc.c)
			if err != nil {
				t.Fatal(err)
			}
			if want, got := tc.wantNum, sum; want != got {
				t.Fatalf("want=%d, got=%d", want, got)
			}

			if want, got := len(tc.wantNums), len(nums); want != got {
				t.Fatalf("want=%d, got=%d", want, got)
			}

			for i := range tc.wantNums {
				if want, got := tc.wantNums[i], nums[i]; want != got {
					t.Errorf("want=%d, got=%d", want, got)
				}
			}
		})
	}
}

func TestDay02Part2(t *testing.T) {
	m, err := convertInputToMatrix("input_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(m); i++ {
		if want, got := len(m[0]), len(m[i]); want != got {
			t.Errorf("input matrix rows are not the same len, want=%d, got=%d", want, got)
		}
	}
	solution, err := doPart2(m)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := 467835, solution; want != got {
		t.Errorf("want=%d, got=%d", want, got)
	}

	m, err = convertInputToMatrix("")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(m); i++ {
		if want, got := len(m[0]), len(m[i]); want != got {
			t.Errorf("input matrix rows are not the same len, want=%d, got=%d", want, got)
		}
	}
	solution, err = doPart2(m)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := 84051670, solution; want != got {
		t.Errorf("want=%d, got=%d", want, got)
	}
}
