package main

import "testing"

func TestSolutionPart2(t *testing.T) {
	t.Run("input_test", func(t *testing.T) {
		out, err := solution("input_test.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(6), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(22289513667691), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
}

func TestPrimeFactor(t *testing.T) {
	testCases := []struct {
		name string
		in   int64
		out  map[int64]int64
	}{
		{
			name: "1",
			in:   1,
			out:  map[int64]int64{},
		},
		{
			name: "2",
			in:   2,
			out:  map[int64]int64{2: 1},
		},
		{
			name: "3",
			in:   3,
			out:  map[int64]int64{3: 1},
		},
		{
			name: "4",
			in:   4,
			out:  map[int64]int64{2: 2},
		},
		{
			name: "30",
			in:   30,
			out:  map[int64]int64{2: 1, 3: 1, 5: 1},
		},
		{
			name: "11",
			in:   11,
			out:  map[int64]int64{11: 1},
		},
		{
			name: "23147",
			in:   23147,
			out:  map[int64]int64{79: 1, 293: 1},
		},
		{
			name: "17287",
			in:   17287,
			out:  map[int64]int64{59: 1, 293: 1},
		},
		{
			name: "21389",
			in:   21389,
			out:  map[int64]int64{73: 1, 293: 1},
		},
		{
			name: "13771",
			in:   13771,
			out:  map[int64]int64{47: 1, 293: 1},
		},
		{
			name: "19631",
			in:   19631,
			out:  map[int64]int64{67: 1, 293: 1},
		},
		{
			name: "20803",
			in:   20803,
			out:  map[int64]int64{71: 1, 293: 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := primeFactoriaztion(tc.in)
			if want, got := len(tc.out), len(out); want != got {
				t.Fatalf("count primes want=%d, got=%d", want, got)
			}
			for k, v := range tc.out {
				gotV, ok := out[k]
				if !ok {
					t.Fatalf("want=%d, key not present", k)
				}
				if want, got := v, gotV; want != got {
					t.Fatalf("key=%d, want=%d, got=%d", k, want, got)
				}
			}
		})
	}
}

func TestLCM(t *testing.T) {
	testCases := []struct {
		name string
		a, b int64
		out  int64
	}{
		{
			name: "1",
			a:    23147,
			b:    17287,
			out:  1365673,
		},
		{
			name: "2",
			a:    1365673,
			b:    21389,
			out:  99694129,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := LCM(tc.a, tc.b)
			if want, got := tc.out, out; want != got {
				t.Fatalf("want=%d, got=%d", want, got)
			}
		})
	}
}
