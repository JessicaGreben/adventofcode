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

var isValid = map[opCode]bool{
	opCodeAdd:       true,
	opCodeMultiply:  true,
	opCodeRead:      true,
	opCodeWrite:     true,
	opCodeTerminate: true,
}

var opCodeInstructionCount = map[opCode]int{
	opCodeAdd:       4,
	opCodeMultiply:  4,
	opCodeRead:      2,
	opCodeWrite:     2,
	opCodeTerminate: 1,
}

type addressingMode int

const (
	// absoluteAddress interprets a parameter as a postion.
	absoluteAddress addressingMode = 0
	// immediateValue interprets a parameter as a value.
	immediateValue addressingMode = 1
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
		op, addressingModes, err := parseOpCodeAddressMode(p.instructions[p.programCounter])
		if err != nil {
			return -1, err
		}
		switch op {
		case opCodeAdd:
			value1, value2, dstIdx := read2In1Out(p.programCounter, p.instructions, addressingModes)
			p.instructions[dstIdx] = value1 + value2
		case opCodeMultiply:
			value1, value2, dstIdx := read2In1Out(p.programCounter, p.instructions, addressingModes)
			p.instructions[dstIdx] = value1 * value2
		case opCodeRead:
			dstIdx := p.instructions[p.programCounter+1]
			p.instructions[dstIdx] = input
		case opCodeWrite:
			dstIdx := p.instructions[p.programCounter+1]
			out := p.instructions[dstIdx]
			if len(addressingModes) >= 1 && addressingModes[0] == immediateValue {
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

func read2In1Out(programCounter int, input []int, addressingModes []addressingMode) (operand1, operand2, op int) {
	operand1, operand2, output := input[programCounter+1], input[programCounter+2], input[programCounter+3]
	if len(addressingModes) == 0 || addressingModes[0] == absoluteAddress {
		operand1 = input[operand1]
	}
	if len(addressingModes) <= 1 || addressingModes[1] == absoluteAddress {
		operand2 = input[operand2]
	}
	return operand1, operand2, output
}

// parseOpCodeAddressMode parses the opCode and the addressing modes from the input.
// The opCode is a two-digit number based only on the ones and tens digit of the value.
// The addressing mode is series of digits starting from the hundreths place.
func parseOpCodeAddressMode(value int) (opCode, []addressingMode, error) {
	opCodeValue := value
	if value >= 100 {
		opCodeValue %= 100
	}

	if ok := isValid[opCode(opCodeValue)]; !ok {
		return opCodeUnknown, []addressingMode{}, fmt.Errorf("%w: %d", errInvalidOpCode, opCodeValue)
	}
	return opCode(opCodeValue), parseAddressingModes(value / 100), nil
}

// parseAddressingModes converts the input integer into a list of its digits,
// then for each digit, converts it to the corresponding addressingMode.
func parseAddressingModes(modes int) []addressingMode {
	digits := intToDigits(modes)
	addressModes := make([]addressingMode, 0, len(digits))
	for _, digit := range digits {
		addressMode := addressingMode(digit)
		if addressMode != absoluteAddress && addressMode != immediateValue {
			return []addressingMode{}
		}
		addressModes = append(addressModes, addressMode)
	}
	return addressModes

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
