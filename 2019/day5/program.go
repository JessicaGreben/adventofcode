package day5

import (
	"golang.org/x/exp/slices"

	"github.com/jessicagreben/adventofcode/2019/pkg/intcode"
)

func TESTDiagnosticProgram(systemID int) []int {
	p := intcode.NewProgram(slices.Clone(AoCInput))
	p.Run(systemID)
	return p.Output()
}

func TESTDiagnosticProgramPart2(systemID int) []int {
	p := intcode.NewProgram(slices.Clone(AoCInput))
	p.Run(systemID)
	return p.Output()
}
