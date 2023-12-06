package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolutionPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  mapping
	}{
		{"1", "50 98 2", mapping{interval{98, 99}, interval{50, 51}}},
		{"2", "52 50 48", mapping{interval{50, 97}, interval{52, 99}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := lineToMapping(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if want, got := tc.want.source.start, out.source.start; want != got {
				t.Errorf("(-want +got):\n%s", cmp.Diff(want, got))
			}
			if want, got := tc.want.source.end, out.source.end; want != got {
				t.Errorf("(-want +got):\n%s", cmp.Diff(want, got))
			}
			if want, got := tc.want.dest.start, out.dest.start; want != got {
				t.Errorf("(-want +got):\n%s", cmp.Diff(want, got))
			}
			if want, got := tc.want.dest.end, out.dest.end; want != got {
				t.Errorf("(-want +got):\n%s", cmp.Diff(want, got))
			}
		})
	}

	t.Run("input_test", func(t *testing.T) {
		out, err := solution("../input_test.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(35), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(51580674), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
}

func TestSolutionForEachLine(t *testing.T) {
	g, err := ForEachLine("../input_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("g: %#v\n", g)

	if want, got := g.seeds, []int64{79, 14, 55, 13}; !cmp.Equal(want, got) {
		t.Errorf("(-want +got):\n%s", cmp.Diff(want, got))
	}

	if want, got := len(g.mappings), 7; !cmp.Equal(want, got) {
		t.Errorf("(-want +got):\n%s", cmp.Diff(want, got))
	}

	if want, got := len(g.mappings[fertilizerCategory]), 3; !cmp.Equal(want, got) {
		t.Errorf("(-want +got):\n%s", cmp.Diff(want, got))
	}

	x, err := g.getLocation(79, soilCategory)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := x, int64(82); want != got {
		t.Errorf("want=%d, got=%d", want, got)
	}
	x, err = g.getLocation(14, soilCategory)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := x, int64(43); want != got {
		t.Errorf("want=%d, got=%d", want, got)
	}
	x, err = g.getLocation(13, soilCategory)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := x, int64(35); want != got {
		t.Errorf("want=%d, got=%d", want, got)
	}
	x, err = g.getLocation(55, soilCategory)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := x, int64(86); want != got {
		t.Errorf("want=%d, got=%d", want, got)
	}
}

func TestFindMappedNum(t *testing.T) {
	m := mapping{interval{50, 97}, interval{52, 99}}
	out, err := m.findMappedNum(53)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := int64(55), out; want != got {
		t.Errorf("want=%d, got=%d", want, got)
	}
	out, err = m.findMappedNum(10)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := int64(-1), out; want != got {
		t.Errorf("want=%d, got=%d", want, got)
	}
	out, err = m.findMappedNum(96)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := int64(98), out; want != got {
		t.Errorf("want=%d, got=%d", want, got)
	}
	out, err = m.findMappedNum(79)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := int64(81), out; want != got {
		t.Errorf("want=%d, got=%d", want, got)
	}

}
