package main

import (
	"testing"
)

func TestSolutionPart2b(t *testing.T) {
	t.Run("input_test", func(t *testing.T) {
		out, err := solution("../input_test.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(46), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		t.Skip()
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(99751240), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
}

func TestConvertInput(t *testing.T) {
	g, err := parseInput("../input_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	for i, x := range g.seedIntervals {
		t.Logf("%d seed interval=%v", i, x)
	}
	for i, x := range g.mappings {
		t.Logf("%s mappings:", i)
		for _, y := range x {
			t.Logf("- %#v\n", y)

		}
	}
}

func TestDo(t *testing.T) {
	g, err := parseInput("../input_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	out := g.getLocationIntervals(interval{79, 92}, 1)
	t.Log(out)
	out = g.getLocationIntervals(interval{55, 67}, 1)
	t.Log(out)
}

func TestConvertInternal(t *testing.T) {
	testCases := []struct {
		name string
		iv   interval
		in   []mapping
		want []interval
	}{
		{
			name: "overlap",
			iv:   interval{79, 92},
			in: []mapping{
				{interval{98, 99}, interval{50, 51}},
				{interval{50, 97}, interval{52, 99}},
			},
			want: []interval{
				{81, 94},
			},
		},
		{
			name: "no overlap",
			iv:   interval{81, 94},
			in: []mapping{
				{interval{15, 51}, interval{0, 36}},
				{interval{52, 53}, interval{37, 38}},
				{interval{0, 14}, interval{39, 53}},
			},
			want: []interval{
				{81, 94},
			},
		},
		{
			name: "two partial overlap",
			iv:   interval{74, 87},
			in: []mapping{
				{interval{45, 63}, interval{81, 99}},
				{interval{77, 99}, interval{45, 67}},
				{interval{64, 76}, interval{68, 80}},
			},
			want: []interval{
				{45, 55},
				{78, 80},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := convertInterval(tc.iv, tc.in)

			if want, got := len(tc.want), len(out); want != got {
				t.Fatalf("want=%d, got=%d", want, got)
			}
			for i := range out {
				if want, got := tc.want[i].start, out[i].start; want != got {
					t.Fatalf("want=%v, got=%v", want, got)
				}
				if want, got := tc.want[i].end, out[i].end; want != got {
					t.Fatalf("want=%v, got=%v", want, got)
				}
			}
		})
	}
}
