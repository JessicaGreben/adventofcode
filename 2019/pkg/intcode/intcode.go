package intcode

import (
	"errors"
	"fmt"
)

type opCode int

const (
	opCodeAdd       opCode = 1
	opCodeMultiply  opCode = 2
	opCodeRead      opCode = 3
	opCodeWrite     opCode = 4
	opCodeTerminate opCode = 99
	opCodeUnknown   opCode = -1
)

var intToOpCode = map[int]opCode{
	1:  opCodeAdd,
	2:  opCodeMultiply,
	3:  opCodeRead,
	4:  opCodeWrite,
	99: opCodeTerminate,
	-1: opCodeUnknown,
}

var opCodeInstructionCount = map[opCode]int{
	opCodeAdd:       4,
	opCodeMultiply:  4,
	opCodeRead:      2,
	opCodeWrite:     2,
	opCodeTerminate: 1,
}

type parameterMode int

const (
	// positionMode interprets a parameter as a postion.
	positionMode parameterMode = 0
	// immediateMode interprets a parameter as a value.
	immediateMode parameterMode = 1
)

var (
	errInvalidOpCode = errors.New("invalid opCode")
)

type Program struct {
	programCounter int
	instructions   []int
	output         []int
}

func NewProgram(input []int) *Program {
	return &Program{
		instructions: input,
		output:       []int{},
	}
}

func (p *Program) Run(input int) (int, error) {
	for p.programCounter < len(p.instructions) {
		op, paramModes, err := parseOpCodeParameterMode(p.instructions[p.programCounter])
		if err != nil {
			return -1, err
		}
		switch op {
		case opCodeAdd:
			value1, value2, dstIdx := read2In1Out(p.programCounter, p.instructions, paramModes)
			p.instructions[dstIdx] = value1 + value2
		case opCodeMultiply:
			value1, value2, dstIdx := read2In1Out(p.programCounter, p.instructions, paramModes)
			p.instructions[dstIdx] = value1 * value2
		case opCodeRead:
			dstIdx := p.instructions[p.programCounter+1]
			p.instructions[dstIdx] = input
		case opCodeWrite:
			dstIdx := p.instructions[p.programCounter+1]
			out := p.instructions[dstIdx]
			if len(paramModes) >= 1 && paramModes[0] == immediateMode {
				out = dstIdx
			}
			p.output = append(p.output, out)
		case opCodeTerminate:
			return p.instructions[0], nil
		default:
			return -1, fmt.Errorf("%w: %d", errInvalidOpCode, op)
		}

		p.programCounter += opCodeInstructionCount[op]
	}
	return p.instructions[0], nil
}

func (p *Program) Output() []int {
	return p.output
}

func read2In1Out(programCounter int, input []int, paramModes []parameterMode) (operand1, operand2, op int) {
	operand1, operand2, output := input[programCounter+1], input[programCounter+2], input[programCounter+3]
	if len(paramModes) == 0 || paramModes[0] == positionMode {
		operand1 = input[operand1]
	}
	if len(paramModes) <= 1 || paramModes[1] == positionMode {
		operand2 = input[operand2]
	}
	return operand1, operand2, output
}

// parseOpCodeParameterMode parses the opCode and the parameter modes from the input.
// opCode is a two-digit number based only on the ones and tens digit of the value.
// parameter mode is series of digits starting from the hundreths place.
func parseOpCodeParameterMode(value int) (opCode, []parameterMode, error) {
	opCodeValue := value
	if value >= 100 {
		opCodeValue %= 100
	}

	opcode, ok := intToOpCode[opCodeValue]
	if !ok {
		return opCodeUnknown, []parameterMode{}, fmt.Errorf("%w: %d", errInvalidOpCode, opCodeValue)
	}
	return opcode, parseParamModes(value / 100), nil
}

// parseParamModes converts the input integer into a list of its digits,
// then for each digit, convert it to the corresponding parameterMode.
func parseParamModes(modes int) []parameterMode {
	digits := intToDigits(modes)
	params := make([]parameterMode, 0, len(digits))
	for _, digit := range digits {
		paramMode := parameterMode(digit)
		if paramMode != positionMode && paramMode != immediateMode {
			return []parameterMode{}
		}
		params = append(params, paramMode)
	}
	return params

}

// intToDigits splits the input integer into a slice of its digits where each digit is an int.
// The least signigicant digit is at index 0 of the output.
func intToDigits(value int) []int {
	digits := []int{}
	n := value
	for n > 0 {
		r := n % 10
		digits = append(digits, r)
		n /= 10
	}
	return digits
}
