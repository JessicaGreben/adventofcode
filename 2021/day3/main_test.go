package main

import (
	"reflect"
	"testing"
)

func TestDecimalToBinary(t *testing.T) {
	testCases := []struct {
		name       string
		input      int
		wantOutput []int
	}{
		{"zero", 0, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{"one", 1, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		{"sixteen", 16, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0}},
		{"nineteen", 19, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1}},
		{"1024", 1024, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		// {"negative", -9, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0}},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			output := decimalToBinary2(tt.input)
			if got, want := output, tt.wantOutput; !reflect.DeepEqual(got, want) {
				t.Fatalf("decimalToBinary2 got %v, want %v", got, want)
			}
			output = decimalToBinary3(tt.input)
			if got, want := output, tt.wantOutput; !reflect.DeepEqual(got, want) {
				t.Fatalf("decimalToBinary3 got %v, want %v", got, want)
			}
		})
	}
}

func BenchmarkDecimalToBinary2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		decimalToBinary2(1230098)
	}
}
func BenchmarkDecimalToBinary3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		decimalToBinary3(1230098)
	}
}
