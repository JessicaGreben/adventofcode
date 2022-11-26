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
	opCodeJIT       opCode = 5
	opCodeJIF       opCode = 6
	opCodeLT        opCode = 7
	opCodeEQ        opCode = 8
	opCodeTerminate opCode = 99
	opCodeUnknown   opCode = -1
)

var isValid = map[opCode]bool{
	opCodeAdd:       true,
	opCodeMultiply:  true,
	opCodeRead:      true,
	opCodeWrite:     true,
	opCodeJIT:       true,
	opCodeJIF:       true,
	opCodeLT:        true,
	opCodeEQ:        true,
	opCodeTerminate: true,
}

var opCodeInstructionCount = map[opCode]int{
	opCodeAdd:       4,
	opCodeMultiply:  4,
	opCodeRead:      2,
	opCodeWrite:     2,
	opCodeJIT:       3,
	opCodeJIF:       3,
	opCodeLT:        4,
	opCodeEQ:        4,
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
	errIndexCount    = errors.New("wrong number of indexes provided")
)

type Program struct {
	programCounter int
	instructions   []int
	output         []int
}

func NewProgram(instructions []int) *Program {
	return &Program{
		instructions: instructions,
		output:       []int{},
	}
}

func (p *Program) Run(input int) error {
	for p.Next() {
		opcode, parameterIndexes, err := p.Parse()
		if err != nil {
			return err
		}
		if opcode == opCodeTerminate {
			return nil
		}
		if opcode == opCodeRead {
			p.instructions[parameterIndexes[0]] = input
		}
		if err := p.Exec(opcode, parameterIndexes); err != nil {
			return err
		}
	}
	return nil
}

func (p *Program) Next() bool {
	return p.programCounter < len(p.instructions)
}

func (p *Program) Parse() (opCode, []int, error) {
	opcode, modes, err := p.parseOpCodeAddressMode()
	if err != nil {
		return opCodeUnknown, []int{}, err
	}
	indexes := p.parseParameterIndexes(opcode, modes)
	return opcode, indexes, err
}

func (p *Program) parseOpCodeAddressMode() (opCode, []addressingMode, error) {
	opCodeValue := p.instructions[p.programCounter]
	if opCodeValue >= 100 {
		opCodeValue %= 100
	}

	if ok := isValid[opCode(opCodeValue)]; !ok {
		return opCodeUnknown, nil, fmt.Errorf("%w: opCode=%d", errInvalidOpCode, opCodeValue)
	}
	modes := parseAddressingModes(p.instructions[p.programCounter] / 100)
	return opCode(opCodeValue), modes, nil
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

func (p *Program) parseParameterIndexes(opcode opCode, modes []addressingMode) []int {
	parameterCount := opCodeInstructionCount[opcode] - 1
	indexes := make([]int, 0, parameterCount)
	for i := 1; i <= parameterCount; i++ {
		mode := absoluteAddress
		if len(modes) > 0 && i-1 < len(modes) {
			mode = modes[i-1]
		}

		idx := p.programCounter + i
		if mode == absoluteAddress {
			indexes = append(indexes, p.instructions[idx])
		}
		if mode == immediateValue {
			indexes = append(indexes, idx)
		}
	}
	return indexes
}

func (p *Program) Exec(op opCode, indexes []int) error {
	switch op {
	case opCodeAdd:
		if len(indexes) < 3 {
			return fmt.Errorf("%w: expected=3, len(indexes)=%d", errIndexCount, len(indexes))
		}
		value1, value2, dstIdx := p.instructions[indexes[0]], p.instructions[indexes[1]], indexes[2]
		p.instructions[dstIdx] = value1 + value2
		p.programCounter += opCodeInstructionCount[op]
	case opCodeMultiply:
		if len(indexes) < 3 {
			return fmt.Errorf("%w: expected=3, len(indexes)=%d", errIndexCount, len(indexes))
		}
		value1, value2, dstIdx := p.instructions[indexes[0]], p.instructions[indexes[1]], indexes[2]
		p.instructions[dstIdx] = value1 * value2
		p.programCounter += opCodeInstructionCount[op]
	case opCodeRead:
		p.programCounter += opCodeInstructionCount[op]
	case opCodeWrite:
		if len(indexes) < 1 {
			return fmt.Errorf("%w: expected=1, len(indexes)=%d", errIndexCount, len(indexes))
		}
		p.output = append(p.output, p.instructions[indexes[0]])
		p.programCounter += opCodeInstructionCount[op]
	case opCodeJIT:
		if len(indexes) < 2 {
			return fmt.Errorf("%w: expected=2, len(indexes)=%d", errIndexCount, len(indexes))
		}
		value1, value2 := p.instructions[indexes[0]], p.instructions[indexes[1]]
		p.programCounter += opCodeInstructionCount[op]
		if value1 != 0 {
			p.programCounter = value2
		}
	case opCodeJIF:
		if len(indexes) < 2 {
			return fmt.Errorf("%w: expected=2, len(indexes)=%d", errIndexCount, len(indexes))
		}
		value1, value2 := p.instructions[indexes[0]], p.instructions[indexes[1]]
		p.programCounter += opCodeInstructionCount[op]
		if value1 == 0 {
			p.programCounter = value2
		}
	case opCodeLT:
		if len(indexes) < 3 {
			return fmt.Errorf("%w: expected=3, len(indexes)=%d", errIndexCount, len(indexes))
		}
		value1, value2, value3 := p.instructions[indexes[0]], p.instructions[indexes[1]], indexes[2]
		p.instructions[value3] = 0
		if value1 < value2 {
			p.instructions[value3] = 1
		}
		p.programCounter += opCodeInstructionCount[op]
	case opCodeEQ:
		if len(indexes) < 3 {
			return fmt.Errorf("%w: expected=3, len(indexes)=%d", errIndexCount, len(indexes))
		}
		value1, value2, value3 := p.instructions[indexes[0]], p.instructions[indexes[1]], indexes[2]
		p.instructions[value3] = 0
		if value1 == value2 {
			p.instructions[value3] = 1
		}
		p.programCounter += opCodeInstructionCount[op]
	default:
		return fmt.Errorf("%w: %d", errInvalidOpCode, op)
	}
	return nil
}

func (p *Program) Output() []int {
	return p.output
}

func (p *Program) Instructions() []int {
	return p.instructions
}
