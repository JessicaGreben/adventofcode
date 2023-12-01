package main

import "testing"

func TestDay01Part1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"1", "4onefive6zsjhzvrjnsfive1", 41},
		{"2", "six5vfb", 55},
		{"3", "6lxpdpdnr", 66},
		{"4", "56cnfourgrkfzxcvpsqd", 56},
		{"5", "oneninepnvtfbbcx98vmttscj64", 94},
		{"6", "cc8one", 88},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := doPart1(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if want, got := tc.want, got; want != got {
				t.Errorf("want=%d, got=%d", want, got)
			}
		})
	}
}

func TestDay01Part2(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"1", "two1nine", 29},
		{"2", "eightwothree", 83},
		{"3", "abcone2threexyz", 13},
		{"4", "xtwone3four", 24},
		{"5", "4nineeightseven2", 42},
		{"6", "zoneight234", 14},
		{"7", "7pqrstsixteen", 76},
		{"8", "sevensevenkpbggfhrhk121", 71},
	}

	var sum int

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			num, err := doPart2(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if want, got := tc.want, num; want != got {
				t.Errorf("want=%d, got=%d", want, got)
			}
			sum += num
		})
	}
	if want, got := 352, sum; want != got {
		t.Errorf("sum want=%d, got=%d", want, got)
	}
}
