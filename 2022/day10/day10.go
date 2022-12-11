package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() int {
	instructions := parseInput("input.txt")
	return runPart1(instructions)
}

func part2() [][]string {
	instructions := parseInput("input.txt")
	return runPart2(instructions)
}

type instructionType string

const (
	addInstruction  instructionType = "addx"
	noopInstruction instructionType = "noop"
)

var cyclesPerInstruction = map[instructionType]int{
	addInstruction:  2,
	noopInstruction: 1,
}

type instruction struct {
	name  instructionType
	value int
}

func parseInput(file string) []instruction {
	fd, err := os.Open(file)
	if err != nil {
		fmt.Println("err Open", err)
	}
	instructions := []instruction{}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		var in instruction

		switch line[0] {
		case string(addInstruction):
			val, err := strconv.Atoi(line[1])
			if err != nil {
				fmt.Println("strconv err", line[1])
			}
			in = instruction{addInstruction, val}
		case string(noopInstruction):
			in = instruction{name: noopInstruction}
		default:
			fmt.Println("not supported type", line[0])
		}

		instructions = append(instructions, in)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("scanner err", err)
	}
	return instructions
}

func runPart1(instructions []instruction) int {
	var signalStrength int
	var registerX int = 1
	var cycle int

	selectCycles := map[int]bool{
		20: true, 60: true, 100: true, 140: true, 180: true, 220: true,
	}
	for _, currInstruction := range instructions {
		for i := 0; i < cyclesPerInstruction[currInstruction.name]; i++ {
			cycle++
			if selectCycles[cycle] {
				signalStrength += (registerX * cycle)
			}
		}
		switch currInstruction.name {
		case addInstruction:
			registerX += currInstruction.value
		default:
			// noop
		}
	}

	return signalStrength
}

func runPart2(instructions []instruction) [][]string {
	var signalStrength int
	var registerX int = 1
	var cycle int

	selectCycles := map[int]bool{
		20: true, 60: true, 100: true, 140: true, 180: true, 220: true,
	}

	finalsRow := make([][]string, 0, 6)
	row := make([]string, 0, 40)
	for _, currInstruction := range instructions {
		for i := 0; i < cyclesPerInstruction[currInstruction.name]; i++ {
			if registerX-1 == (cycle%40) || registerX == (cycle%40) || registerX+1 == (cycle%40) {
				row = append(row, "#")
			} else {
				row = append(row, ".")
			}
			cycle++
			if selectCycles[cycle] {
				signalStrength += (registerX * cycle)
			}
			if cycle%40 == 0 {
				finalsRow = append(finalsRow, row)
				row = make([]string, 0, 40)
			}
		}
		switch currInstruction.name {
		case addInstruction:
			registerX += currInstruction.value
		default:
			// noop
		}
	}

	return finalsRow
}
