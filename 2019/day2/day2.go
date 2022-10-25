package day2

import (
	"errors"
	"fmt"
)

const (
	opcodeAdd       = 1
	opcodeMultiply  = 2
	opcodeTerminate = 99
)

var opcodeInstructionCount = map[int]int{
	opcodeAdd:       4,
	opcodeMultiply:  4,
	opcodeTerminate: 1,
}

var (
	errInvalidOpcode = errors.New("invalid opcode")
	errInvalidLen    = errors.New("invalid length")
	errNounRange     = errors.New("noun out of [0,99] range")
	errVerbRange     = errors.New("verb out of [0,99] range")
)

func Part1(input []int) (int, error) {
	const (
		noun = 12
		verb = 2
	)
	if err := restore(noun, verb, input); err != nil {
		return -1, err
	}
	return IntcodeProgram(input)
}

func restore(noun, verb int, input []int) error {
	const minInputLen = 3
	if len(input) < minInputLen {
		return fmt.Errorf("%w: minimum length is %d, got %d", errInvalidLen, minInputLen, len(input))
	}
	if noun < 0 || noun > 99 {
		return fmt.Errorf("%w: %d", errNounRange, noun)
	}
	if verb < 0 || verb > 99 {
		return fmt.Errorf("%w: %d", errVerbRange, verb)
	}
	input[1] = noun
	input[2] = verb
	return nil
}

func IntcodeProgram(input []int) (int, error) {
	var programCounter int
	for programCounter < len(input) {
		op := input[programCounter]
		switch {
		case op == opcodeAdd || op == opcodeMultiply:
			inputPtr1, inputPtr2, outputPtr := input[programCounter+1], input[programCounter+2], programCounter+3
			value1, value2, dstIdx := input[inputPtr1], input[inputPtr2], input[outputPtr]
			var out int
			if op == opcodeAdd {
				out = value1 + value2
				programCounter += opcodeInstructionCount[opcodeAdd]
			}
			if op == opcodeMultiply {
				out = value1 * value2
				programCounter += opcodeInstructionCount[opcodeMultiply]
			}
			input[dstIdx] = out
		case op == opcodeTerminate:
			return input[0], nil
		default:
			return -1, fmt.Errorf("%w: %d", errInvalidOpcode, op)
		}
	}
	return input[0], nil
}

func Part2(input []int) (int, int, error) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			cp := makeCopy(input)
			restore(noun, verb, cp)
			x, err := IntcodeProgram(cp)
			if err != nil {
				return -1, -1, err
			}
			if x == 19690720 {
				return noun, verb, nil
			}
		}
	}
	return -1, -1, nil
}

func makeCopy(arr []int) []int {
	cp := make([]int, len(arr))
	copy(cp, arr)
	return cp
}
