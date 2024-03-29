package main

import (
	"fmt"

	"github.com/jessicagreben/adventofcode/pkg/input"
)

func parseInput(file string) []matrix {
	result := []matrix{}
	m, err := input.ConvertToMatrix(file)
	if err != nil {
		fmt.Println("err parsing:", err)
		return nil
	}

	nextMatrix := matrix{data: [][]string{}}
	for _, row := range m {
		if len(row) == 0 {
			result = append(result, nextMatrix)
			nextMatrix = matrix{data: [][]string{}}

			continue
		}
		nextMatrix.data = append(nextMatrix.data, row)
	}
	return append(result, nextMatrix)
}
