package main

import (
	"math"
	"strings"

	"github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int, error) {
	return input.ForEachLine(file, processLine)
}

func processLine(line string) (int, error) {
	nums := strings.Split(line, "|")
	winNums, gameNums := strings.Split(nums[0], " "), strings.Split(nums[1], " ")

	winNumsMap := map[string]bool{}
	for _, num := range winNums {
		winNumsMap[num] = true
	}

	var countMatchWinNums int = 0
	for _, n := range gameNums {
		if n == "" {
			continue
		}
		if _, ok := winNumsMap[n]; ok {
			countMatchWinNums++
		}
	}

	if countMatchWinNums == 0 {
		return 0, nil
	}
	return int(math.Pow(2, float64(countMatchWinNums)-1)), nil
}
