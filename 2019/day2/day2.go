package day2

import (
	"errors"
	"fmt"

	"golang.org/x/exp/slices"
)

type opCode int

const (
	opCodeAdd       opCode = 1
	opCodeMultiply  opCode = 2
	opCodeTerminate opCode = 99
)

var opCodeInstructionCount = map[opCode]int{
	opCodeAdd:       4,
	opCodeMultiply:  4,
	opCodeTerminate: 1,
}

var (
	errInvalidOpCode = errors.New("invalid opCode")
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
	return RunProgram(input)
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

func RunProgram(input []int) (int, error) {
	var programCounter int
	for programCounter < len(input) {
		op := opCode(input[programCounter])
		switch op {
		case opCodeAdd:
			value1, value2, dstIdx := getValues(programCounter, input)
			input[dstIdx] = value1 + value2
		case opCodeMultiply:
			value1, value2, dstIdx := getValues(programCounter, input)
			input[dstIdx] = value1 * value2
		case opCodeTerminate:
			return input[0], nil
		default:
			return -1, fmt.Errorf("%w: %d", errInvalidOpCode, op)
		}

		programCounter += opCodeInstructionCount[op]
	}
	return input[0], nil
}

func getValues(programCounter int, input []int) (operand1, operand2, op int) {
	operandPtr1, operandPtr2, outputPtr := input[programCounter+1], input[programCounter+2], programCounter+3
	return input[operandPtr1], input[operandPtr2], input[outputPtr]
}

func Part2(input []int) (int, int, error) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			cp := slices.Clone(input)
			restore(noun, verb, cp)
			x, err := RunProgram(cp)
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
