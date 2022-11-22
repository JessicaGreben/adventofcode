package day5

import (
	"github.com/jessicagreben/adventofcode/2019/pkg/intcode"
)

func TESTDiagnosticProgram(systemID int) []int {
	p := intcode.NewProgram(AoCInput)
	p.Run(systemID)
	return p.Output()
}
