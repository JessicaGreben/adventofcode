package intcode

import (
	"errors"
	"testing"
)

func TestParseOpCodeParameterMode(t *testing.T) {
	testCases := []struct {
		name      string
		input     int
		wantCode  opCode
		wantModes []addressingMode
		err       error
	}{
		{"1", 3, opCodeRead, []addressingMode{}, nil},
		{"2", 1, opCodeAdd, []addressingMode{}, nil},
		{"3", 1100, opCodeUnknown, []addressingMode{}, errInvalidOpCode},
		{"4", 104, opCodeWrite, []addressingMode{immediateValue}, nil},
		{"5", 1102, opCodeMultiply, []addressingMode{immediateValue, immediateValue}, nil},
		{"6", 6, opCodeUnknown, []addressingMode{}, errInvalidOpCode},
		{"7", 99, opCodeTerminate, []addressingMode{}, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotCode, gotModes, gotErr := parseOpCodeAddressMode(tc.input)
			if !errors.Is(gotErr, tc.err) {
				t.Errorf("want: %v, got: %v", tc.err, gotErr)
			}
			if want, got := len(tc.wantModes), len(gotModes); want != got {
				t.Fatalf("want: %d, got: %d", want, got)
			}
			if want, got := tc.wantCode, gotCode; want != got {
				t.Errorf("want: %d, got: %d", want, got)
			}
			for i, mode := range gotModes {
				if want, got := tc.wantModes[i], mode; want != got {
					t.Errorf("want: %d, got: %d", want, got)
				}
			}
		})
	}
}

func TestIntToDigits(t *testing.T) {
	testCases := []struct {
		name  string
		input int
		want  []int
	}{
		{"1", 0, []int{}},
		{"2", 1, []int{1}},
		{"3", 12, []int{2, 1}},
		{"4", 123, []int{3, 2, 1}},
		{"5", 1101, []int{1, 0, 1, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := intToDigits(tc.input)
			if want, got := len(tc.want), len(output); want != got {
				t.Fatalf("want: %d, got: %d", want, got)
			}

			for i, val := range output {
				if want, got := tc.want[i], val; want != got {
					t.Fatalf("want: %d, got: %d", want, got)
				}
			}
		})
	}
}
