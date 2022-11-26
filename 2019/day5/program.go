package day5

import (
	"fmt"

	"golang.org/x/exp/slices"

	"github.com/jessicagreben/adventofcode/2019/pkg/intcode"
)

func TESTDiagnosticProgram(systemID int) []int {
	p := intcode.NewProgram(slices.Clone(AoCInput))
	if err := p.Run(systemID); err != nil {
		fmt.Println("err intcode.Run", err)
	}
	return p.Output()
}

func TESTDiagnosticProgramPart2(systemID int) []int {
	p := intcode.NewProgram(slices.Clone(AoCInput))
	if err := p.Run(systemID); err != nil {
		fmt.Println("err intcode.Run", err)
	}

	return p.Output()
}
