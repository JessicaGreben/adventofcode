package main

import "testing"

func TestSolutionPart1(t *testing.T) {
	t.Run("input_test", func(t *testing.T) {
		out, err := solution("input_test.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := 2, out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := 23147, out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})
}

func TestSolutionParseInput(t *testing.T) {
	t.Run("input_test", func(t *testing.T) {
		p, n, err := ForEachLine("input_test.txt")
		if err != nil {
			t.Fatal(err)
		}

		if want, got := 2, len(p); want != got {
			t.Fatalf("want=%d, got=%d", want, got)
		}
		want := []string{"R", "L"}
		for i := range p {
			if want, got := want[i], p[i]; want != got {
				t.Fatalf("want=%s, got=%s", want, got)
			}
		}
		if want, got := 7, len(n); want != got {
			t.Fatalf("want=%d, got=%d", want, got)
		}
		wantNodes := map[string][]string{
			"AAA": []string{"BBB", "CCC"},
			"BBB": []string{"DDD", "EEE"},
			"CCC": []string{"ZZZ", "GGG"},
			"DDD": []string{"DDD", "DDD"},
			"EEE": []string{"EEE", "EEE"},
			"GGG": []string{"GGG", "GGG"},
			"ZZZ": []string{"ZZZ", "ZZZ"},
		}
		for k := range n {
			if want, got := len(wantNodes[k]), len(n[k]); want != got {
				t.Fatalf("want=%d, got=%d", want, got)
			}
			want := wantNodes[k]
			for i := range want {
				if want, got := want[i], n[k][i]; want != got {
					t.Fatalf("want=%s, got=%s", want, got)
				}
			}
		}

	})
}
