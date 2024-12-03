package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	lines, err := input.GetLines(file)
	if err != nil {
		return -1, err
	}

	regex := `mul\(\d+,\d+\)`
	re := regexp.MustCompile(regex)
	var sum int64
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		for _, m := range matches {
			// each part should look like this mul(x,y)
			parts := strings.FieldsFunc(m, splitOn)
			if len(parts) != 3 {
				return -1, fmt.Errorf("wrong len parts, want=3, got=%v", len(parts))
			}
			xInt, err := strconv.Atoi(parts[1])
			if err != nil {
				return -1, err
			}
			yInt, err := strconv.Atoi(parts[2])
			if err != nil {
				return -1, err
			}
			sum += int64(xInt) * int64(yInt)
		}
	}
	return sum, nil
}

func splitOn(r rune) bool {
	return r == '(' || r == ',' || r == ')'
}
