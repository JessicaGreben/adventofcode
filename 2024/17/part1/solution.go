package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func solution(registerA, registerB, registerC int64, program []int) (string, error) {
	output := []int64{}
	var instructionPointer int

	for instructionPointer+1 < len(program) {
		opcode := program[instructionPointer]
		operand := program[instructionPointer+1]

		instructionPointer += 2
		switch o := opcode; o {
		case 0:
			numerator := registerA
			denominator := math.Pow(float64(2), float64(getComboValue(operand, registerA, registerB, registerC)))
			registerA = numerator / int64(denominator)
		case 1:
			registerB ^= int64(operand)
		case 2:
			registerB = getComboValue(operand, registerA, registerB, registerC) % 8
		case 3:
			if registerA == 0 {
				continue
			}
			instructionPointer = operand
		case 4:
			registerB ^= registerC
		case 5:
			output = append(output, getComboValue(operand, registerA, registerB, registerC)%8)
		case 6:
			numerator := registerA
			denominator := math.Pow(float64(2), float64(getComboValue(operand, registerA, registerB, registerC)))
			registerB = numerator / int64(denominator)
		case 7:
			numerator := registerA
			denominator := math.Pow(float64(2), float64(getComboValue(operand, registerA, registerB, registerC)))
			registerC = numerator / int64(denominator)
		}
	}

	var outpuStr []string
	for _, out := range output {
		outpuStr = append(outpuStr, strconv.FormatInt(out, 10))
	}
	return strings.Join(outpuStr, ","), nil
}

// Combo operands 0 through 3 represent literal values 0 through 3.
// Combo operand 4 represents the value of register A.
// Combo operand 5 represents the value of register B.
// Combo operand 6 represents the value of register C.
// Combo operand 7 is reserved and will not appear in valid programs.
func getComboValue(operand int, registerA, registerB, registerC int64) int64 {
	switch o := operand; o {
	case 0:
		return int64(operand)
	case 1:
		return int64(operand)
	case 2:
		return int64(operand)
	case 3:
		return int64(operand)
	case 4:
		return registerA
	case 5:
		return registerB
	case 6:
		return registerC
	default:
		fmt.Println("err unsupported combo", o)
		return -1
	}
}
