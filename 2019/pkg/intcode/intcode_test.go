package intcode

import (
	"testing"
)

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

func TestParse(t *testing.T) {
	testCases := []struct {
		name         string
		instructions []int
		opcode       opCode
		output       []int
	}{
		{"1a opCode 3", []int{3, 9}, opCodeRead, []int{9}},
		{"1b opCode 3", []int{103, 10}, opCodeRead, []int{1}},
		{"2a opcode 8", []int{8, 9, 10, 9}, opCodeEQ, []int{9, 10, 9}},
		{"2b opcode 8", []int{108, 9, 10, 300}, opCodeEQ, []int{1, 10, 300}},
		{"2c opcode 8", []int{1108, 9, 10, 300}, opCodeEQ, []int{1, 2, 300}},
		{"2d opcode 8", []int{11108, 9, 10, 300}, opCodeEQ, []int{1, 2, 3}},
		{"3a opcode 4", []int{4, 10}, opCodeWrite, []int{10}},
		{"3b opcode 4", []int{104, 1}, opCodeWrite, []int{1}},
		{"4a opCode 99", []int{99}, opCodeTerminate, []int{}},
		{"5a opcode 6", []int{6, 9, 10}, opCodeJIF, []int{9, 10}},
		{"5b opcode 6", []int{106, 9, 10}, opCodeJIF, []int{1, 10}},
		{"5c opcode 6", []int{1106, 9, 10}, opCodeJIF, []int{1, 2}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewProgram(tc.instructions)
			opcode, parameterIndexes, err := p.ParseOp()
			if err != nil {
				t.Error(err)
			}
			if want, got := tc.opcode, opcode; want != got {
				t.Errorf("opcode want %v, got %v", want, got)
			}
			if want, got := tc.output, parameterIndexes; len(want) != len(got) {
				t.Fatalf("len want: %v, got: %v", want, got)
			}
			for i, val := range parameterIndexes {
				if want, got := tc.output[i], val; want != got {
					t.Errorf("want: %v, got: %v", want, got)
				}
			}
		})
	}
}

func TestProgram(t *testing.T) {
	testCases := []struct {
		name         string
		instructions []int
		input        int
		output       []int
	}{
		{"1a opCode 8, input LT", []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, 5, []int{0}},
		{"1b opcode 8, input equal", []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, 8, []int{1}},
		{"1c opcode 8, input GT", []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, 10, []int{0}},
		{"2a opCode 7, input LT", []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, 5, []int{1}},
		{"2b opcode 7, input equal", []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, 8, []int{0}},
		{"2c opcode 7, input GT", []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, 10, []int{0}},
		{"3a opCode 8, input LT", []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, 5, []int{0}},
		{"3b opcode 8, input equal", []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, 8, []int{1}},
		{"3c opcode 8, input GT", []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, 10, []int{0}},
		{"4a opCode 7, input LT", []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, 5, []int{1}},
		{"4b opcode 7, input equal", []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, 8, []int{0}},
		{"4c opcode 7, input GT", []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, 10, []int{0}},
		{"5a opcode 6, input 10", []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, 10, []int{1}},
		{"5b opcode 6, input 0", []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, 0, []int{0}},
		{"6a opcode 5, input non-zero", []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, 10, []int{1}},
		{"6b opcode 5, input zero", []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, 0, []int{0}},
		{"7a larger example LT", []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 5, []int{999}},
		{"7b larger example equal", []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 8, []int{1000}},
		{"7c larger example GT", []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 10, []int{1001}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewProgram(tc.instructions)
			p.Run(tc.input)
			if want, got := len(tc.output), len(p.Output()); want != got {
				t.Fatalf("len want: %d, got: %d", want, got)
			}
			for i, val := range p.Output() {
				if want, got := tc.output[i], val; want != got {
					t.Fatalf("want: %d, got: %d", want, got)
				}
			}
		})
	}
}
