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

	noJump = -1
)

var (
	errInvalidOpCode = errors.New("invalid opCode")
	errIndexCount    = errors.New("wrong number of indexes provided")
)

type Program struct {
	programCounter int
	memory         []int
	output         []int
}

func NewProgram(memory []int) *Program {
	return &Program{
		memory: memory,
		output: []int{},
	}
}

func (p *Program) Run(input int) error {
	for {
		opcode, parameterIndexes, err := p.ParseOp()
		if err != nil {
			return err
		}
		if opcode == opCodeTerminate {
			return nil
		}
		if opcode == opCodeRead {
			p.memory[parameterIndexes[0]] = input
		}
		jumpTo, err := p.Exec(opcode, parameterIndexes)
		if err != nil {
			return err
		}

		p.programCounter += opCodeInstructionCount[opcode]
		if jumpTo != noJump {
			p.programCounter = jumpTo
		}
	}
}

func (p *Program) ParseOp() (opCode, []int, error) {
	opcode, modes, err := p.parseOpCodeAddressMode()
	if err != nil {
		return opCodeUnknown, []int{}, err
	}
	indexes := p.parseParameterIndexes(opcode, modes)
	return opcode, indexes, err
}

func (p *Program) parseOpCodeAddressMode() (opCode, []addressingMode, error) {
	// Parse opcode converts the instruction to an opcode.
	instruction := p.memory[p.programCounter]
	opCodeValue := instruction
	if opCodeValue >= 100 {
		opCodeValue %= 100
	}
	if ok := isValid[opCode(opCodeValue)]; !ok {
		return opCodeUnknown, nil, fmt.Errorf("%w: opCode=%d", errInvalidOpCode, opCodeValue)
	}

	// Parse address modes converts the input integer into a list of its digits.
	// Each digit is converted it to the corresponding addressingMode.
	modeValue := instruction / 100
	digit := iterateDigits(modeValue)
	addressModes := []addressingMode{}
	for digit.next() {
		addressModes = append(addressModes, addressingMode(digit.value()))
	}
	return opCode(opCodeValue), addressModes, nil
}

type digit struct {
	n int
}

func iterateDigits(value int) *digit {
	return &digit{
		n: value,
	}
}

func (d *digit) next() bool {
	return d.n > 0
}

func (d *digit) value() int {
	r := d.n % 10
	d.n /= 10
	return r
}

// parseParameterIndexes converts the addressing modes to the index of the location of the parameters
// for the opcode instruction.
func (p *Program) parseParameterIndexes(opcode opCode, modes []addressingMode) []int {
	parameterCount := opCodeInstructionCount[opcode] - 1
	indexes := make([]int, 0, parameterCount)
	for modeIdx, parameterIdx := 0, p.programCounter+1; modeIdx < parameterCount; modeIdx, parameterIdx = modeIdx+1, parameterIdx+1 {
		mode := absoluteAddress
		if len(modes) > 0 && modeIdx < len(modes) {
			mode = modes[modeIdx]
		}

		if mode == absoluteAddress {
			indexes = append(indexes, p.memory[parameterIdx])
		}
		if mode == immediateValue {
			indexes = append(indexes, parameterIdx)
		}
	}
	return indexes
}

func (p *Program) Exec(op opCode, indexes []int) (int, error) {
	if len(indexes) < opCodeInstructionCount[op]-1 {
		return noJump, fmt.Errorf("%w: expected=%d, len(indexes)=%d", errIndexCount, opCodeInstructionCount[op]-1, len(indexes))
	}
	switch op {
	case opCodeAdd:
		operand1, operand2, dstIdx := p.memory[indexes[0]], p.memory[indexes[1]], indexes[2]
		p.memory[dstIdx] = operand1 + operand2
	case opCodeMultiply:
		operand1, operand2, dstIdx := p.memory[indexes[0]], p.memory[indexes[1]], indexes[2]
		p.memory[dstIdx] = operand1 * operand2
	case opCodeRead:
		// no op
	case opCodeWrite:
		p.output = append(p.output, p.memory[indexes[0]])
	case opCodeJIT:
		jumpCondition, jumpTo := p.memory[indexes[0]], p.memory[indexes[1]]
		if jumpCondition != 0 {
			return jumpTo, nil
		}
	case opCodeJIF:
		jumpCondition, jumpTo := p.memory[indexes[0]], p.memory[indexes[1]]
		if jumpCondition == 0 {
			return jumpTo, nil
		}
	case opCodeLT:
		operand1, operand2, dstIdx := p.memory[indexes[0]], p.memory[indexes[1]], indexes[2]
		p.memory[dstIdx] = 0
		if operand1 < operand2 {
			p.memory[dstIdx] = 1
		}
	case opCodeEQ:
		operand1, operand2, dstIdx := p.memory[indexes[0]], p.memory[indexes[1]], indexes[2]
		p.memory[dstIdx] = 0
		if operand1 == operand2 {
			p.memory[dstIdx] = 1
		}
	default:
		return noJump, fmt.Errorf("%w: opcode=%d", errInvalidOpCode, op)
	}
	return noJump, nil
}

func (p *Program) Output() []int {
	return p.output
}

func (p *Program) Memory() []int {
	return p.memory
}
